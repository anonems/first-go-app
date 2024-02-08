package routes

import (
	"hairdresser-app/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	router.POST("/login", controllers.Login())
	//router.GET("/logout", controllers.Logout())
}
