package rest

import (
	someresource "async_worker/internal/someResource"

	"github.com/gin-gonic/gin"
)

func BuildRoutes(router *gin.Engine, deps *someresource.SomeResourceDependencies) {
	router.GET("/ping", ping)
	someresource.BuildRoutes(router, deps)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
