package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zanovru/gin-rest-api/internal/services"
)

type Routing struct {
	services *services.Services
}

func NewRouting(services *services.Services) *Routing {
	return &Routing{services: services}
}

func (r *Routing) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(logRequests)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/auth", r.auth)
		v1.POST("/register", r.register)
	}

	return router
}
