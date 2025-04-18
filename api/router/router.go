package router

import (
	"net/http"

	"github.com/Wefdzen/medods/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter.
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// support swagger
	r.StaticFS("/swagger-docs", http.Dir("./api/swagger"))
	url := ginSwagger.URL("/swagger-docs/swagger.yml")
	r.GET("/ui-swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// taks endpoints
	r.POST("/api/token", handler.IssueTokensHandler())
	r.POST("/api/refresh", handler.RefreshTokensHandler())

	return r
}
