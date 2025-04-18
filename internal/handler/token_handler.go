package handler

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/Wefdzen/medods/internal/db"
	"github.com/Wefdzen/medods/internal/service"
	"github.com/gin-gonic/gin"
)

// тело post
type RequestBody struct {
	GUID string `json:"guid"`
}

func IssueTokensHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if nginx => X-Forwarded-For
		ipOfClientTmp := c.Request.RemoteAddr
		ipOfClient, err := service.ParseIPv6(ipOfClientTmp)
		if err != nil {
			log.Fatal(err)
			return
		}

		var requestBody RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		// validate guid
		if err := service.ValidateGuid(requestBody.GUID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "didn't validate",
			})
			return
		}

		// gen unic code
		unicCode := service.GenUnicCode()

		// Generate tokens
		accessToken, err := service.GenerateAccessToken(requestBody.GUID, ipOfClient, unicCode)
		if err != nil {
			log.Println(err)
			return
		}
		refreshToken, liveOfRefToken, err := service.GenerateRefreshToken(requestBody.GUID, ipOfClient, unicCode)
		if err != nil {
			log.Println(err)
			return
		}

		// work with database , refTok to bcrypt
		userRepo := db.NewGormUserRepository()

		hashRefreshToken, err := service.HashString(refreshToken)
		if err != nil {
			log.Fatal(err)
			return
		}
		user := db.User{
			Guid:             requestBody.GUID,
			RefreshTokenHash: hashRefreshToken,
			IpClient:         ipOfClient,
			UnicCode:         unicCode,
			LiveToken:        fmt.Sprintf("%v", liveOfRefToken),
		}
		// check uniq
		if !db.CheckUniqGuid(userRepo, user.Guid) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Your guid isn't uniq",
			})
			return
		}
		// add to database
		err = db.AddRecord(userRepo, &user)
		if err != nil {
			log.Fatal(err)
			return
		}

		// set refreshToken to base64
		encodedRefToken := base64.StdEncoding.EncodeToString([]byte(refreshToken))

		// setCookie
		c.Copy().SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("accessToken", accessToken, 3600*24*30, "", "", false, true)
		c.SetCookie("refreshToken", encodedRefToken, 3600*24*30, "", "", false, true)

		c.Status(http.StatusOK)
	}
}
