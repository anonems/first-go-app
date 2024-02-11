package routes

import (
	"hairdresser-app/controllers"
	"github.com/gin-gonic/gin"
)

func HairdresserRoute(router *gin.Engine) {

	//admin action
	router.POST("/hairdresser", controllers.CreateHairdresser())
	router.POST("/hairdresser/edit/:hairdresserId", controllers.EditHairdresser())
	router.GET("/hairdresser/delete/:hairdresserId", controllers.DeleteHairdresser())

	//user action

}
