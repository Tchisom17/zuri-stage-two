package server

import (
	"github.com/gin-gonic/gin"
	"zuri-stage-one/internal/adapters/api"
	"zuri-stage-one/internal/core/middlewares"
)

func SetupRouter(handler *api.HTTPHandler) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORS())

	r := router.Group("/api/v1")
	{
		r.GET("/", handler.Get)
	}

	return router
}
