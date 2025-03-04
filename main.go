package main

import (
	"github.com/gin-gonic/gin"
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	err := r.Run()
	if err != nil {
		return
	}
}
