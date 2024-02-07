package main

import (
	"backend/configs"
	"backend/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)

	router.Run("localhost:5000")
	fmt.Println("Hello word!")
}