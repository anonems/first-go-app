package routes

import (
	"hairdresser-app/controllers"

	"github.com/gin-gonic/gin"
)

func TemplateRoute(router *gin.Engine) {

	//user page
	router.GET("/", controllers.HomePageTemplate())
	router.GET("/menu", controllers.MainMenuTemplate())
	router.GET("/myProfile", controllers.MyProfileTemplate())
	router.GET("/myCompany", controllers.MyCompanyTemplate())
	router.GET("/myAppointments", controllers.MyAppointmentsTemplate())
	router.GET("/adminMenu", controllers.AdminMenuTemplate())
	router.GET("/appointmentType", controllers.AppointmentTypesTemplate())
	router.GET("/appointmentType/edit/:typeId", controllers.EditAppointmentTypesTemplate())
	router.GET("/hairdresser", controllers.HairdresserTemplate())
	router.GET("/hairdresser/edit/:hairdresserId", controllers.EditHairdresserTemplate())
	router.GET("/appointment/admin", controllers.AdminAppointmentList())
	router.GET("/appointment/admin/edit/:appointmentId", controllers.EditAdminAppointment())

}
