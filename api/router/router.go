package router

import (
	"github.com/Wefdzen/medods/internal/handler"
	"github.com/gin-gonic/gin"
)

// SetupRouter.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/token", handler.IssueTokensHandler())
	r.POST("/api/refresh", handler.RefreshTokensHandler())

	return r
}
