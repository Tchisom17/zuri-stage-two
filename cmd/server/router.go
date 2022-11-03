package server

import (
	"github.com/gin-gonic/gin"
	"zuri-stage-one/internal/adapters/api"
	"zuri-stage-one/internal/core/middleware"
)

func SetupRouter(handler *api.HTTPHandler) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())

	r := router.Group("/api/v1")
	{
		r.GET("/", handler.Get)
		r.POST("/calc", handler.Calculate)
	}

	return router
}
