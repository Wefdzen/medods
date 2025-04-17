package handler

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/Wefdzen/medods/internal/service"
	"github.com/gin-gonic/gin"
)

// тело post
type RequestBody struct {
	GUID string `json:"guid"`
}

func IssueTokensHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipOfClient := c.Request.RemoteAddr

		var requestBody RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		//validate guid
		if err := service.ValidateGuid(requestBody.GUID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "didn't validate",
			})
			return
		}

		//gen unic code
		unicCode := service.GenUnicCode()

		//Generate tokens
		accessToken, err := service.GenerateAccessToken(requestBody.GUID, ipOfClient, unicCode)
		if err != nil {
			log.Println(err)
			return
		}
		refreshToken, err := service.GenerateRefreshToken(requestBody.GUID, ipOfClient, unicCode)
		if err != nil {
			log.Println(err)
			return
		}

		//set refreshToken to base64
		encodedRefToken := base64.StdEncoding.EncodeToString([]byte(refreshToken))

		//setCookie
		c.Copy().SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("accessToken", accessToken, 3600*24*30, "", "", false, true)
		c.SetCookie("refreshToken", encodedRefToken, 3600*24*30, "", "", false, true)

		//TODO work with database , refTok to bcrypt

		c.Status(http.StatusOK)
	}
}
