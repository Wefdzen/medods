package handler

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/Wefdzen/medods/internal/db"
	"github.com/Wefdzen/medods/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RefreshTokensHandler() gin.HandlerFunc {
	// Get couple of tokens
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("accessToken")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "you haven't cookie accessToken",
			})
			return
		}
		refreshTokenBase64, err := c.Cookie("refreshToken")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "you haven't cookie refreshToken",
			})
			return
		}
		// TODO сейчас оно работает что пофик сдох или нет токен
		claimsAccessToken, _ := service.ParseToken(c, accessToken)

		// Decode refreshToken
		refreshToken, err := base64.StdEncoding.DecodeString(refreshTokenBase64)
		if err != nil {
			log.Fatal("RefreshToken was not base64")
			return
		}
		claimsRefreshToken, _ := service.ParseToken(c, string(refreshToken))

		// Check unicCode of couple of tokens
		if claimsAccessToken["unicCode"].(string) != claimsRefreshToken["unicCode"].(string) {
			log.Fatal("unicCode of tokens isn't equals")
			return
		}

		userRepo := db.NewGormUserRepository()

		user, err := db.GetRecord(userRepo, claimsRefreshToken["sub"].(string))
		if err != nil {
			log.Fatal(err)
			return
		}

		hash := sha256.Sum256(refreshToken)
		// Check equals of our refToken with refToken from db
		if err := bcrypt.CompareHashAndPassword([]byte(user.RefreshTokenHash), hash[:]); err != nil {
			log.Fatal(err)
			return
		}
		// Проверяем айпи откуда был запрос с айпи из бд
		ipOfClientTmp := c.Request.RemoteAddr
		ipOfClient, err := service.ParseIPv6(ipOfClientTmp)
		if err != nil {
			log.Fatal(err)
			return
		}

		if user.IpClient != ipOfClient {
			c.JSON(http.StatusBadRequest, gin.H{"error": "IP текущего пк не равен айпи того кто создавал "})
			// make smtp схерали у чела есть токены но он на другом пк:#
			// You can change here stubEmailService
			err := service.SendWarningEmail(&service.StubEmailService{}, claimsAccessToken["sub"].(string), claimsAccessToken["IpClient"].(string), ipOfClient)
			if err != nil {
				log.Fatal(err)
				return
			}
			return
		}

		// gen unic code
		unicCode := service.GenUnicCode()

		// Generate tokens
		newAccessToken, err := service.GenerateAccessToken(claimsAccessToken["sub"].(string), ipOfClient, unicCode)
		if err != nil {
			log.Println(err)
			return
		}
		newRefreshToken, newLiveToken, err := service.GenerateRefreshToken(claimsAccessToken["sub"].(string), ipOfClient, unicCode)
		if err != nil {
			log.Println(err)
			return
		}

		// update in db
		hashNew, _ := service.HashString(newRefreshToken)
		db.UpdateReftokenLiveTokenUnicCode(userRepo, claimsAccessToken["sub"].(string), hashNew, fmt.Sprintf("%v", newLiveToken), unicCode)

		// set refreshToken to base64
		encodedRefToken := base64.StdEncoding.EncodeToString([]byte(newRefreshToken))

		// setCookie
		c.Copy().SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("accessToken", newAccessToken, 3600*24*30, "", "", false, true)
		c.SetCookie("refreshToken", encodedRefToken, 3600*24*30, "", "", false, true)

		c.Status(http.StatusOK)
	}
}
