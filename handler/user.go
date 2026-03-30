package handler

import (
	"net/http"
	"user-service/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
)

func PublicHandler(c *gin.Context) {
	message := service.GetPublicMessage()

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
func ProfileHandler(c *gin.Context) {
	userClaims, _ := c.Get("user")

	claims, ok := userClaims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Invalid claims format",
		})
		return
	}

	result := service.GetProfile(claims)

	c.JSON(http.StatusOK, result)
}