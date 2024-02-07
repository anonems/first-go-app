package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func HairCompany(router *gin.Engine) {
	router.POST("/hairCompany", controllers.CreateHairCompany())
	router.GET("/hairCompany/:companyId", controllers.GetHairCompany())
	router.PUT("/hairCompany/:companyId", controllers.EditHairCompany())
	router.DELETE("/hairCompany/:companyId", controllers.DeleteHairCompany())
	router.GET("hairCompanies", controllers.GetAllHairCompanies())
}