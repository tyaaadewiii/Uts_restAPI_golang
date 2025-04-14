package controllers

import (
	"github.com/gin-gonic/gin"
	"uts_restapi_golang/config"
	"uts_restapi_golang/models"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&user)
	c.JSON(201, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&user)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User tidak ditemukan"})
		return
	}

	config.DB.Delete(&user)
	c.JSON(200, gin.H{"message": "User berhasil dihapus"})
}
