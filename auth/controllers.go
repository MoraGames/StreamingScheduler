package main

import (
	"encoding/json"
	"fmt"
	"github.com/MoraGames/StreamingScheduler/auth/internal/jwt"
	"github.com/MoraGames/StreamingScheduler/auth/internal/mail"
	"github.com/MoraGames/StreamingScheduler/auth/internal/utils"
	"net/http"
	"os"
	"strconv"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetIP(r)

	var params struct {
		Email    string
		Password string
	}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get user info
	user, err := GetUserByEmail(params.Email)
	if err != nil {
		utils.PrintErr(w, "Incorrect username or password")
		return
	}

	// verify password and email
	if (user.Password != params.Password) || (user.Email != params.Email) {
		utils.PrintErr(w, "Incorrect username or password")
		return
	}

	// verify account active
	if !user.Enabled {
		http.Error(w, `{"code": 401, "msg": "Account inactive!"}`, http.StatusUnauthorized)
		return
	}

	// get access token expiration
	accessExp, err := strconv.Atoi(os.Getenv("JWT_AT_EXP"))
	if err != nil {
		log.Error("Error to get access token expiration: " + err.Error())
		utils.PrintInternalErr(w)
		return
	}

	// get refresh token expiration
	refreshExp, err := strconv.Atoi(os.Getenv("JWT_RT_EXP"))
	if err != nil {
		log.Error("Error to get refresh token expiration: " + err.Error())
		utils.PrintInternalErr(w)
		return
	}

	log.Infoln(user)

	tokenPair := jwt.NewJWTokenPair(
		accessExp,
		refreshExp,
	)
	tokenPair.Access.Obj.Email = params.Email
	tokenPair.Access.Obj.Iss = os.Getenv("JWT_ISS")
	tokenPair.Access.Obj.Iat = time.Now().Unix()
	tokenPair.Access.Obj.Company = os.Getenv("JWT_COMPANY")
	tokenPair.Access.Obj.Permission = user.Permissions.ToString()
	tokenPair.Refresh.Obj.Email = params.Email
	tokenPair.Refresh.Obj.RefreshId = utils.GenerateID()

	err = tokenPair.GenerateTokenPair(os.Getenv("JWT_AT_PWD"), os.Getenv("JWT_RT_PWD"))
	if err != nil {
		log.Error("General", ip, "ApiLogin", "Error to generate token pair: "+err.Error())
		utils.PrintInternalErr(w)
		return
	}

	//Set Cookies
	err = utils.SetCookies(w, tokenPair.Refresh.Token, os.Getenv("JWT_ISS"), os.Getenv("COOKIE_SECRET"))
	if err != nil {
		log.Error("General", ip, "ApiLogin", "Error to set cookies: "+err.Error())
		utils.PrintInternalErr(w)
		return
	}

	// Add new refresh token to database
	err = jwt.AddToDB(dbConn, tokenPair.Refresh.Token, tokenPair.Refresh.Obj.Exp, user.Id)
	if err != nil {
		log.Error("General", ip, "ApiLogin", "Database operation error: "+err.Error())
		utils.PrintInternalErr(w)
		return
	}

	data2, err := json.Marshal(map[string]string{
		"AccessToken": tokenPair.Access.Token,
		"Expiration":  fmt.Sprint(tokenPair.Access.Obj.Exp),
	})
	if err != nil {
		log.Info("General", ip, "ApiLogin", "Error to create AccessToken JSON: "+err.Error())
		utils.PrintInternalErr(w)
		return
	}

	w.Write(data2)
}

func register(w http.ResponseWriter, r *http.Request) {

	var u User

	// Declare a new MusicData struct.
	user := struct {
		Email          string `json:"email"`
		Username       string `json:"username"`
		Password       string `json:"password"`
		ProfilePicture string `json:"profilePicture,omitempty"`
	}{}

	ip := utils.GetIP(r)

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Error("Error to decode registration body: " + err.Error())
		utils.PrintErr(w, err.Error())
		return
	}

	u.Email = user.Email
	u.Username = user.Username
	u.Password = user.Password
	u.ProfilePicture = user.ProfilePicture
	u.Permissions = jwt.Permission("u")

	err = u.IsValid()
	if err != nil {
		utils.PrintErr(w, err.Error())
		return
	}

	// Check user exist
	exist, err := u.Exist()
	if err != nil {
		log.Error("Error to check exist of user: " + err.Error())
		utils.PrintInternalErr(w)
		return
	}

	if exist {
		utils.PrintErr(w, "A user with this email already exists")
		return
	}

	//Add in the database
	lastId, err := u.NewUser()
	if err != nil {
		log.Error("General", ip, "ApiSignUp", "AddNewUser error: "+err.Error())
		utils.PrintInternalErr(w)
		return
	}

	// Get refresh token expire
	refreshExp, err := strconv.Atoi(os.Getenv("JWT_RT_EXP"))
	if err != nil {
		log.Error("Error to get the refresh token expiration:", err.Error())
		return
	}

	//Generate refresh token
	refToken := jwt.NewRefreshToken(refreshExp)
	refToken.RefreshId = utils.GenerateID()
	refToken.Email = u.Email

	type Conferma struct {
		Username string
		Link     string
		Login    string
	}

	//Save refresh token in the database
	err = jwt.AddToDB(dbConn, refToken.RefreshId, refToken.Exp, lastId)
	if err != nil {
		log.Error("Database error to add refresh token: " + err.Error())
		utils.PrintInternalErr(w)
		return
	}

	c := Conferma{
		Username: u.Username,
		Link: fmt.Sprintf(
			"https://%s:%s/api/v1/confirm?email=%s&id=%s",
			os.Getenv("HOSTNAME"),
			os.Getenv("PORT"),
			u.Email,
			refToken.RefreshId,
		),
		Login: "https://" + os.Getenv("HOSTNAME") + ":" + os.Getenv("PORT") + "/api/v1/login",
	}

	//Send mails
	err = mail.SendEmail(
		os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
		os.Getenv("EMAIL_ADDR"),
		os.Getenv("EMAIL_PWD"),
		u.Email,
		"CONFERMA REGISTRAZIONE",
		"template/email/registrazione.txt",
		c,
	)
	if err != nil {
		log.Error("Error to send mail: " + err.Error())
		return
	}

	w.Write([]byte(`{"id": ` + strconv.Itoa(int(lastId)) + "}"))
}

func verify(w http.ResponseWriter, r *http.Request) {

	// Get client ip
	ip := utils.GetIP(r)

	// Get client params
	params := []utils.ParamsInfo{
		{Key: "id", Required: true},
		{Key: "email", Required: true},
	}

	p, err := utils.GetParams(params, r)
	if err != nil {
		log.Error("General", ip, "ApiConfirmSignup", "Error to get params: "+err.Error())
		utils.PrintInternalErr(w)
		return
	}

	// Get user info
	user, err := GetUserByEmail(p["email"].(string))
	if err != nil {
		log.Error("Error to get user info:", err.Error())
		utils.PrintInternalErr(w)
		return
	}

	// Verify token
	ok, err := jwt.VerifyRefreshToken(dbConn, user.Id, p["id"].(string))
	if err != nil {
		log.Error(err)
		utils.PrintInternalErr(w)
		return
	}

	if !ok {
		log.Error(p["email"].(string), ip, "ApiConfirmSignup", "Warning to verify refresh token: Token not valid")
		utils.PrintErr(w, "Token not valid!")
		return
	}

	//Remove old refresh token
	err = jwt.RemoveRefreshToken(dbConn, p["id"].(string))
	if err != nil {
		log.Error("Error to remove expired refresh token")
		utils.PrintInternalErr(w)
		return
	}

	//Set user to active
	err = user.Active()
	if err != nil {
		log.Error("Error to enable the user:", err.Error())
		utils.PrintInternalErr(w)
		return
	}

	w.Write([]byte("{\"state\": \"success\"}"))
}
