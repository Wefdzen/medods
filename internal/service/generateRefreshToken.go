package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateAccessToken unicCode для того что бы пара былак взаимосвязанны
// liveToken надо для бд я же немогу из hash of refToken достать
// это поле
func GenerateRefreshToken(guid, ipClient, unicCode string) (string, int64, error) {
	liveToken := time.Now().Add(time.Hour * 120).Unix() // 5day
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub":       guid,
		"liveToken": liveToken,
		"ipClient":  ipClient,
		"unicCode":  unicCode,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("super_secret_key")))
	if err != nil {
		return "", 0, err
	}

	return tokenString, liveToken, nil
}
