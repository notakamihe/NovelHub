package utils

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id uint) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 1440).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return token, nil
}

func SliceContainsIgnoreCase(arr []string, str string) bool {
	for _, v := range arr {
		if strings.ToLower(v) == strings.ToLower(str) {
			return true
		}
	}

	return false
}
