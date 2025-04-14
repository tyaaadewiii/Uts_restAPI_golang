package routes

import (
	"github.com/gin-gonic/gin"
	"uts_restapi_golang/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// USER
	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	// KOSA KATA
	router.GET("/kosakata", controllers.GetKosaKata)
	router.POST("/kosakata", controllers.CreateKosaKata)
	router.PUT("/kosakata/:id", controllers.UpdateKosaKata)
	router.DELETE("/kosakata/:id", controllers.DeleteKosaKata)
}
