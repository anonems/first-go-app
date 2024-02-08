package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
//var validate = validator.New()

func HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Signin / Signup",
		})
	}
}