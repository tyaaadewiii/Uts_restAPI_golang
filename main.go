package main

import (
	"github.com/gin-gonic/gin"
	"uts_restapi_golang/config"
	"uts_restapi_golang/models"
	"uts_restapi_golang/routes"
)

func main() {
	r := gin.Default()
	config.ConnectDB()

	// Optional: Auto migrate (tidak bikin ulang kalau tabel sudah ada)
	config.DB.AutoMigrate(&models.User{}, &models.KosaKata{})

	routes.SetupRoutes(r)

	r.Run() // default di :8080
}
