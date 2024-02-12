package routes

import (
	"hairdresser-app/controllers"
	"github.com/gin-gonic/gin"
)

func AppointmentTypeRoute(router *gin.Engine) {

	//admin action
	router.POST("/appointmentType", controllers.CreateAppointmentType())
	router.POST("/appointmentType/edit/:typeId", controllers.EditAppointmentType())
	router.GET("/appointmentType/delete/:typeId", controllers.DeleteAppointmentType())

	//user action

}
