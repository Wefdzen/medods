package service

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(c *gin.Context, tokenJWT string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenJWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("super_secret_key")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) >= claims["liveToken"].(float64) { // если accessToken умер
			return claims, nil // Я проверял при меньшем времени и оно оказывалось меньше следовательно nil return and error
		} else {
			// TODO пока проверка времени бесполезная time < liveToken
			return claims, nil // если время еще не вышло и accessToken валиден
		}
	}
	return nil, errors.New("Can't parse Caims of token")
}
