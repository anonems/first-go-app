package routes

import (
	"hairdresser-app/controllers"
	"github.com/gin-gonic/gin"
)

func AppointmentRoute(router *gin.Engine) {

	//admin action
	router.POST("/appointment/admin", controllers.CreateAppointment())
	router.POST("/appointment/admin/edit/:appointmentId", controllers.EditAppointment())
	router.GET("/appointment/admin/delete/:appointmentId", controllers.DeleteAppointment())

}
