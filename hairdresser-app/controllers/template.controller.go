package controllers

import (
	"fmt"
	"hairdresser-app/models"
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

func HairCompaniesPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := store.Get(c.Request, "session")
		var user = &models.User{}
		val := session.Values["user"]
		var ok bool
		fmt.Println(user)
		if user, ok = val.(*models.User); !ok {
			fmt.Println("error")
			c.HTML(http.StatusForbidden, "index.html", gin.H{
				"title": "Hair Companies List",
			})
		}
		c.HTML(http.StatusOK, "hairCompanies.html", gin.H{
			"title": "error",
		})
	}
}