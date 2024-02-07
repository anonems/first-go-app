package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func Appointment(router *gin.Engine) {
	router.POST("/appointment", controllers.TakeAppointment())
	router.PUT("/appointment/:appointmentId", controllers.EditAppointment())
	router.DELETE("/appointment/:appointmentId", controllers.CancelAppointment())
}