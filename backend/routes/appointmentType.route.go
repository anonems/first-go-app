package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func AppointmentTypeRoute(router *gin.Engine) {

	//admin action
	router.POST("/appointmentType", controllers.CreateAppointmentType())
	router.PUT("/appointmentType/:appointmentTypeId", controllers.EditAppointmentType())
	router.GET("/appointmentType/:appointmentTypeId", controllers.GetAllAppointmentsType())

	//user action

	
}