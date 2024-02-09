package routes

import (
	"hairdresser-app/controllers"
	"github.com/gin-gonic/gin"
)

func TemplateRoute(router *gin.Engine) {

	//user page
	router.GET("/", controllers.HomePage())
	router.GET("/menu", controllers.MainMenu())
	router.GET("/myProfile", controllers.MyProfile())
	router.GET("/myCompany", controllers.MyCompany())

}