package handler

import (
	"net/http"
	"user-service/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func GetAllUser(c *gin.Context) {
	userClaims, _ := c.Get("user")

	_, ok := userClaims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Invalid claims format",
		})
		return
	}

	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil ambil semua user",
		"data":    users,
	})
}