package main

import (
	"hairdresser-app/configs"
	//"hairdresser-app/models"
	//"fmt"
	"hairdresser-app/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type EmailRequestBody struct {
	Email string
}

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge:   60 * 60 * 24}) // expire in a day
	router.Use(sessions.Sessions("mysession", store))

	//run database
	configs.ConnectDB()

	//routes
	routes.TemplateRoute(router)
	routes.AuthRoute(router)
	routes.UserRoute(router)

	router.Run(":5000")
}
