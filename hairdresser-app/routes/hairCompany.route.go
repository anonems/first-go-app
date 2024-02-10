package routes

import (
	"hairdresser-app/controllers"

	"github.com/gin-gonic/gin"
)

func HairCompanyRoute(router *gin.Engine) {

	//admin action
	router.POST("/hairCompany", controllers.CreateHairCompany())
	router.POST("/hairCompany/edit/:companyId", controllers.EditHairCompany())
	router.DELETE("/hairCompany/:companyId", controllers.DeleteHairCompany())

	//user action
	router.GET("/hairCompany/:companyId", controllers.GetHairCompany())
	router.GET("hairCompanies", controllers.GetAllHairCompanies())

}
