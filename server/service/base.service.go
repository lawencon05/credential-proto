package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CatchError for catch with param error
func CatchError(e *error) {
	if err := recover(); err != nil {
		*e = fmt.Errorf("%v", err)
	}
}

const secret = "secret"

// GenerateToken with username as param
func GenerateToken(username string) (t string, e error) {
	defer CatchError(&e)

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
