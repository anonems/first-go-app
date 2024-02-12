package main

import (
	"hairdresser-app/configs"
	"hairdresser-app/routes"
	"github.com/gin-gonic/gin"
)

type EmailRequestBody struct {
	Email string
}

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	//run database
	configs.ConnectDB()

	//routes
	routes.TemplateRoute(router)
	routes.AuthRoute(router)
	routes.UserRoute(router)
	routes.HairCompanyRoute(router)
	routes.AppointmentTypeRoute(router)
	routes.HairdresserRoute(router)
	routes.AppointmentRoute(router)

	router.Run(":5000")
}
