package controllers

import (
	"context"
	"fmt"
	"hairdresser-app/models"
	//"hairdresser-app/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

//var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
//var validate = validator.New()

func HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user = &models.User{}
		session, _ := store.Get(c.Request, "session")
		val := session.Values["user"]
		var ok bool
		if user, ok = val.(*models.User); !ok {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Signin / Signup",
			})
		} else {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var userHairCompany models.UserHairCompany
			defer cancel()

			adminAction := "Manage My Company"

			err := userHairCompanyCollection.FindOne(ctx, bson.M{"userId": user.ID}).Decode(&userHairCompany)

			if err != nil {
				adminAction = "Create Company"

			}

			c.HTML(http.StatusForbidden, "menu.html", gin.H{
				"title":       "Main Menu",
				"adminAction": adminAction,
			})
		}

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

func MainMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := store.Get(c.Request, "session")
		var user = &models.User{}
		val := session.Values["user"]
		var ok bool
		fmt.Println(user)
		if user, ok = val.(*models.User); !ok {
			c.Redirect(http.StatusForbidden, "/")
		} else {
			c.HTML(http.StatusOK, "menu.html", gin.H{
				"title": "Main Menu",
			})
		}

	}
}

func MyProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user = &models.User{}
		session, _ := store.Get(c.Request, "session")
		val := session.Values["user"]
		var ok bool
		if user, ok = val.(*models.User); !ok {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Signin / Signup",
			})
		} else {
			c.HTML(http.StatusOK, "myProfile.html", gin.H{
				"title":     "My Profile",
				"userId":    user.ID.Hex(),
				"firstName": user.FirstName,
				"lastName":  user.LastName,
				"age":       user.Age,
				"email":     user.Email,
				"gender":    user.Gender,
			})

		}
	}
}

func MyCompany() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user = &models.User{}
		session, _ := store.Get(c.Request, "session")
		val := session.Values["user"]
		var ok bool
		if user, ok = val.(*models.User); !ok {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Signin / Signup",
			})
		} else {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var hairCompany models.HairCompany
			var userHairCompany models.UserHairCompany
			defer cancel()

			formTitle := "Update Company"

			err := userHairCompanyCollection.FindOne(ctx, bson.M{"userId": user.ID}).Decode(&userHairCompany)
			rootUrl := "/hairCompany/edit/" + userHairCompany.ID.Hex()
			if err != nil {
				formTitle = "Create Company"
				rootUrl = "/hairCompany"

			}
			hairCompanyCollection.FindOne(ctx, bson.M{"_id": userHairCompany.HairCompanyId}).Decode(&hairCompany)
			c.HTML(http.StatusOK, "myCompany.html", gin.H{
				"title":         "My Company",
				"rootUrl":       rootUrl,
				"formTitle":     formTitle,
				"name":          hairCompany.Name,
				"siren":         hairCompany.SIREN,
				"line1":         hairCompany.Address.Line1,
				"line2":         hairCompany.Address.Line2,
				"postalCode":    hairCompany.Address.PostalCode,
				"city":          hairCompany.Address.City,
				"country":       hairCompany.Address.Country,
			})
		}
	}
}
