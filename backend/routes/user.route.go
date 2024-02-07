package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	//admin action
	router.GET("/user/:userId", controllers.GetUser())

	//user action
	router.POST("/user", controllers.CreateUser())
	router.PUT("/user/:userId", controllers.EditUser())
	
}