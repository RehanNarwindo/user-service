package main

import (
	"user-service/handler"
	"user-service/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/public", handler.PublicHandler)

	r.GET("/profile", middleware.AuthMiddleware(), handler.ProfileHandler)

	r.Run(":3001")
}