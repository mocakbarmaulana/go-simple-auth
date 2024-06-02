package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-simple-auth/controllers"
	"github.com/go-simple-auth/middleware"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	routes := r.Group("/api/v1")
	{
		routes.POST("/login", controllers.Login)
		routes.POST("/register", controllers.Register)
	}

	protectedRoutes := r.Group("/api/v1")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		protectedRoutes.GET("/user", controllers.GetUser)
	}

	r.Run()
}
