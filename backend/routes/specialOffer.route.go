package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func SpecialOfferRoute(router *gin.Engine) {

	//admin action
	router.POST("/specialOffer", controllers.CreateSpecialOffer())
	router.PUT("/specialOffer", controllers.EditSpecialOffer())

	//user action
	router.GET("/specialOffer", controllers.GetAllSpecialOffer())
	
}