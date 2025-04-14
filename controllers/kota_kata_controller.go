package controllers

import (
	"github.com/gin-gonic/gin"
	"uts_restapi_golang/config"
	"uts_restapi_golang/models"
)

func GetKosaKata(c *gin.Context) {
	var kosa []models.KosaKata
	config.DB.Find(&kosa)
	c.JSON(200, kosa)
}

func CreateKosaKata(c *gin.Context) {
	var kosa models.KosaKata
	if err := c.ShouldBindJSON(&kosa); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&kosa)
	c.JSON(201, kosa)
}

func UpdateKosaKata(c *gin.Context) {
	id := c.Param("id")
	var kosa models.KosaKata

	if err := config.DB.First(&kosa, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Kosa kata tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&kosa); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&kosa)
	c.JSON(200, kosa)
}

func DeleteKosaKata(c *gin.Context) {
	id := c.Param("id")
	var kosa models.KosaKata

	if err := config.DB.First(&kosa, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Kosa kata tidak ditemukan"})
		return
	}

	config.DB.Delete(&kosa)
	c.JSON(200, gin.H{"message": "Kosa kata berhasil dihapus"})
}
