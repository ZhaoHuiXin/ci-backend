package web

import (
	"github.com/gin-gonic/gin"
	"ci-backend/internal/web/handlers"
)


func InitRouter() *gin.Engine{
	router := gin.Default()
	v1 := router.Group("/v1")
	// v1 router group
	{
		v1.GET("/hello", handlers.Hello)
		//v1.GET("/test", handlers.Test)
	}
	return router
}
