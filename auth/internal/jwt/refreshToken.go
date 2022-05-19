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

// NewRefreshToken is a function that creates a new refresh token
func NewRefreshToken(addExpire int) *JWTRefreshMetadata {
	return &JWTRefreshMetadata{Exp: time.Now().Add(time.Minute * time.Duration(addExpire)).Unix()}
}

// GenerateToken is a function that generates a new refresh token
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

// VerifyRefreshToken is a function that verify the refresh token
func VerifyRefreshToken(db *sql.DB, idUser int64, idRefresh string) (bool, error) {

	var exists bool
	err := db.QueryRow("SELECT exists (SELECT * FROM RefreshTokens WHERE id = ? && user = ?)", idRefresh, idUser).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.New("error checking if row exists " + err.Error())
	}

	return exists, nil
}

// ExtractRefreshMetadata is a function that extracts the refresh token metadata
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

// AddToDB is a function that adds the refresh token in the database
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

// RemoveRefreshToken is a function that removes the user refresh token from the database
func RemoveRefreshToken(db *sql.DB, token string) error {

	// prepare the insert query
	stmt, err := db.Prepare("DELETE FROM RefreshTokens WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(token)
	if err != nil {
		return err
	}

	return nil
}

// RemoveExpiredRefreshToken is a function that removes the expired user refresh token from the database
func RemoveExpiredRefreshToken(db *sql.DB) error {

	// prepare the insert query
	stmt, err := db.Prepare("DELETE FROM RefreshTokens WHERE exp < ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now().Unix())
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
