package jwt

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/MoraGames/StreamingScheduler/auth/internal/utils"
	"github.com/form3tech-oss/jwt-go"
	"strconv"
	"time"
)

type JWTRefreshMetadata struct {
	RefreshId string
	Email     string
	Exp       int64
}

func NewRefreshToken(addExpire int) *JWTRefreshMetadata {
	return &JWTRefreshMetadata{Exp: time.Now().Add(time.Minute * time.Duration(addExpire)).Unix()}
}

func (jm *JWTRefreshMetadata) GenerateToken(refreshPass string) (string, error) {

	if err := jm.check(); err != nil {
		return "", err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refreshId"] = jm.RefreshId
	rtClaims["email"] = jm.Email
	rtClaims["exp"] = jm.Exp
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := rt.SignedString([]byte(refreshPass))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func VerifyRefreshToken(db *sql.DB, email, idRefresh string) bool {

	// TODO: Implements verify refresh token with new db
	/*
		_, err := db.Client.C.Collection("User").Doc(email).
			Collection("RefreshToken").Doc(idRefresh).
			Get(db.Client.Ctx)
		if err != nil {
			return false
		}
	*/

	return true
}

func ExtractRefreshMetadata(tokenString, secret string) (*JWTRefreshMetadata, error) {

	token, err := VerifyToken(tokenString, secret)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {

		fmt.Println(claims)

		refreshId, ok := claims["refreshId"].(string)
		if !ok {
			return nil, errors.New("RefreshId error!")
		}

		exp, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["exp"]), 10, 64)
		if err != nil {
			return nil, err
		}

		email, ok := claims["email"].(string)
		if !ok {
			return nil, errors.New("Email error!")
		}

		return &JWTRefreshMetadata{
			RefreshId: refreshId,
			Email:     email,
			Exp:       exp,
		}, nil
	}
	return nil, err
}

func AddToDB(db *sql.DB, token string, exp int64, userId int64) error {

	// prepare the insert query
	stmt, err := db.Prepare("INSERT INTO RefreshTokens (id, exp, user) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(token, exp, userId)
	if err != nil {
		return err
	}

	return nil
}

func (jm *JWTRefreshMetadata) check() error {

	if !utils.IsEmailValid(jm.Email) {
		return errors.New("Email not valid")
	}

	if jm.Exp < time.Now().Unix() {
		return errors.New("Expiration not valid")
	}

	return nil
}
