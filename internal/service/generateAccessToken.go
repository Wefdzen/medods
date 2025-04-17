package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateAccessToken unicCode для того что бы пара былак взаимосвязанны
func GenerateAccessToken(guid, ipClient, unicCode string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub":       guid,
		"liveToken": time.Now().Add(time.Minute * 20).Unix(),
		"ipClient":  ipClient,
		"unicCode":  unicCode,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("super_secret_key")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
