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
	"go.mongodb.org/mongo-driver/bson/primitive"
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
			hidden := ""

			err := userHairCompanyCollection.FindOne(ctx, bson.M{"userId": user.ID}).Decode(&userHairCompany)

			if err != nil {
				adminAction = "Create Company"
				hidden = "hidden"
			}

			c.HTML(http.StatusForbidden, "mainMenu.html", gin.H{
				"title":       "Main Menu",
				"adminAction": adminAction,
				"hidden":      hidden,
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
			c.HTML(http.StatusOK, "mainMenu.html", gin.H{
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
			rootUrl := "/hairCompany/edit/" + userHairCompany.HairCompanyId.Hex()
			if err != nil {
				formTitle = "Create Company"
				rootUrl = "/hairCompany"

			}
			hairCompanyCollection.FindOne(ctx, bson.M{"_id": userHairCompany.HairCompanyId}).Decode(&hairCompany)
			c.HTML(http.StatusOK, "myCompany.html", gin.H{
				"title":      "My Company",
				"rootUrl":    rootUrl,
				"formTitle":  formTitle,
				"name":       hairCompany.Name,
				"siren":      hairCompany.SIREN,
				"line1":      hairCompany.Address.Line1,
				"line2":      hairCompany.Address.Line2,
				"postalCode": hairCompany.Address.PostalCode,
				"city":       hairCompany.Address.City,
				"country":    hairCompany.Address.Country,
			})
		}
	}
}

func MyAppointments() gin.HandlerFunc {
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
				"title":      "My Company",
				"rootUrl":    rootUrl,
				"formTitle":  formTitle,
				"name":       hairCompany.Name,
				"siren":      hairCompany.SIREN,
				"line1":      hairCompany.Address.Line1,
				"line2":      hairCompany.Address.Line2,
				"postalCode": hairCompany.Address.PostalCode,
				"city":       hairCompany.Address.City,
				"country":    hairCompany.Address.Country,
			})
		}
	}
}

func AdminMenu() gin.HandlerFunc {
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

			err := userHairCompanyCollection.FindOne(ctx, bson.M{"userId": user.ID}).Decode(&userHairCompany)
			if err != nil {
				c.Redirect(http.StatusFound, "/myCompany")

			}
			hairCompanyCollection.FindOne(ctx, bson.M{"_id": userHairCompany.HairCompanyId}).Decode(&hairCompany)
			c.HTML(http.StatusOK, "adminMenu.html", gin.H{
				"title": "Admin Portal",
			})
		}
	}
}

func AppointmentTypes() gin.HandlerFunc {
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

			err := userHairCompanyCollection.FindOne(ctx, bson.M{"userId": user.ID}).Decode(&userHairCompany)
			if err != nil {
				c.Redirect(http.StatusFound, "/myCompany")

			}

			appointmentTypeCollection.Find(ctx, bson.M{"_id": userHairCompany.HairCompanyId})

			filter := bson.D{primitive.E{Key: "hairCompanyId", Value: userHairCompany.HairCompanyId}}

			// Retrieves documents that match the query filer
			cursor, err := appointmentTypeCollection.Find(context.TODO(), filter)
			if err != nil {
				panic(err)
			}
			var results []models.AppointmentType
			if err = cursor.All(context.TODO(), &results); err != nil {
				panic(err)
			}

			c.HTML(http.StatusOK, "typesList.html", gin.H{
				"title": "Manage Appointment Types",
				"types": results,
			})
		}
	}
}

func EditAppointmentTypes() gin.HandlerFunc {
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
			typeId := c.Param("typeId")
			objId, _ := primitive.ObjectIDFromHex(typeId)
			var userHairCompany models.UserHairCompany
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var appointmentType models.AppointmentType
			defer cancel()
			err := userHairCompanyCollection.FindOne(ctx, bson.M{"userId": user.ID}).Decode(&userHairCompany)
			rootUrl := "/appointmentType"
			formTitle := "Create Type"
			if err != nil {
				c.Redirect(http.StatusFound, "/myCompany")
			}
			appointmentTypeCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&appointmentType)
			if appointmentType.ID.Hex() == typeId {
				rootUrl = "/appointmentType/edit/" + typeId
				formTitle = "Update Type"

			}
			c.HTML(http.StatusOK, "editType.html", gin.H{
				"title":          "Manage Appointment Types",
				"rootUrl":        rootUrl,
				"formTitle":      formTitle,
				"name":           appointmentType.Name,
				"description":    appointmentType.Description,
				"duration":       appointmentType.Duration,
			})
		}
	}
}
