package routes

import (
	"hairdresser-app/controllers"

	"github.com/gin-gonic/gin"
)

func TemplateRoute(router *gin.Engine) {

	//admin pages
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

	//public pages
	router.GET("public/country", controllers.Country())
	router.GET("public/:country/city", controllers.City())
	router.GET("public/company/:country/:city", controllers.Company())
	router.GET("public/type/:companyId", controllers.Type())
	router.GET("public/hairdresser/:typeId", controllers.Hairdresser())
	router.GET("public/appointment/:typeId/:hairdresserId", controllers.Appointment())

}
