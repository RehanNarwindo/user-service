package main

import (
	"user-service/config"
	"user-service/handler"
	"user-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB() 

	r.GET("/public", handler.PublicHandler)

	r.GET("/profile", middleware.AuthMiddleware(), handler.ProfileHandler)

	r.GET("/users", middleware.AuthMiddleware(), handler.GetAllUser)


	r.Run(":3001")
}