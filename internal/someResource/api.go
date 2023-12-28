package someresource

import (
	"async_worker/internal/someResource/app"
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

type SomeResourceRouter struct {
	service *app.Service
}

func BuildRoutes(router *gin.Engine, deps *SomeResourceDependencies) {
	srr := &SomeResourceRouter{
		service: app.NewService(deps.JobProcessor),
	}

	someResourcePaths := router.Group("/some-resource")
	someResourcePaths.POST("", srr.handleSomeResourceRequest)
}

func (srr *SomeResourceRouter) handleSomeResourceRequest(c *gin.Context) {
	srr.service.HandleServiceRequest(c, fmt.Sprintf("user %d", rand.Intn(1000)))
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
