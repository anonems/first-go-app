package routes

import (
	"hairdresser-app/controllers"
	"github.com/gin-gonic/gin"
)

func TemplateRoute(router *gin.Engine) {

	//user page
	router.GET("/", controllers.HomePage())
	router.GET("/list", controllers.HairCompaniesPage())

}