package controllers

import (
	"context"
	"fmt"
	"hairdresser-app/models"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
//var validate = validator.New()

func HomePageTemplate() gin.HandlerFunc {
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

func HairCompaniesPageTemplate() gin.HandlerFunc {
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

func MainMenuTemplate() gin.HandlerFunc {
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

func MyProfileTemplate() gin.HandlerFunc {
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

func MyCompanyTemplate() gin.HandlerFunc {
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

func MyAppointmentsTemplate() gin.HandlerFunc {
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
			c.HTML(http.StatusOK, "myAppointments.html", gin.H{
				"title": "My Appointments",
			})
		}
	}
}

func AdminMenuTemplate() gin.HandlerFunc {
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

func AppointmentTypesTemplate() gin.HandlerFunc {
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
				"list":  results,
			})
		}
	}
}

func EditAppointmentTypesTemplate() gin.HandlerFunc {
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
				"title":       "Manage Hairdresser",
				"rootUrl":     rootUrl,
				"formTitle":   formTitle,
				"name":        appointmentType.Name,
				"description": appointmentType.Description,
				"duration":    appointmentType.Duration,
			})
		}
	}
}

func HairdresserTemplate() gin.HandlerFunc {
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

			filter := bson.D{primitive.E{Key: "hairCompanyId", Value: userHairCompany.HairCompanyId}}

			// Retrieves documents that match the query filer
			cursor, err := hairdresserCollection.Find(context.TODO(), filter)
			if err != nil {
				panic(err)
			}
			var results []models.Hairdresser
			if err = cursor.All(context.TODO(), &results); err != nil {
				panic(err)
			}

			c.HTML(http.StatusOK, "hairdresserList.html", gin.H{
				"title": "Manage Hairdresser Types",
				"list":  results,
			})
		}
	}
}

func EditHairdresserTemplate() gin.HandlerFunc {
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
			hairdresserId := c.Param("hairdresserId")
			objId, _ := primitive.ObjectIDFromHex(hairdresserId)
			var userHairCompany models.UserHairCompany
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var hairdresser models.Hairdresser
			defer cancel()
			err := userHairCompanyCollection.FindOne(ctx, bson.M{"userId": user.ID}).Decode(&userHairCompany)
			rootUrl := "/hairdresser"
			formTitle := "Create Hairdresser"
			if err != nil {
				c.Redirect(http.StatusFound, "/myCompany")
			}
			hairdresserCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&hairdresser)
			if hairdresser.ID.Hex() == hairdresserId {
				rootUrl = "/hairdresser/edit/" + hairdresserId
				formTitle = "Update Hairdresser"

			}

			filter := bson.D{primitive.E{Key: "hairCompanyId", Value: userHairCompany.HairCompanyId}}
			cursor, err := appointmentTypeCollection.Find(context.TODO(), filter)
			if err != nil {
				panic(err)
			}
			var results []models.AppointmentType
			if err = cursor.All(context.TODO(), &results); err != nil {
				panic(err)
			}
			c.HTML(http.StatusOK, "editHairdresser.html", gin.H{
				"title":     "Manage Hairdresser",
				"rootUrl":   rootUrl,
				"formTitle": formTitle,
				"firstName": hairdresser.FirstName,
				"lastName":  hairdresser.LastName,
				"typeId":    hairdresser.TypeId,
				"list": results,
			})
		}
	}
}

func AdminAppointmentList() gin.HandlerFunc {
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

			filter := bson.D{primitive.E{Key: "hairCompanyId", Value: userHairCompany.HairCompanyId}}

			// Retrieves documents that match the query filer
			cursor, err := appointmentCollection.Find(context.TODO(), filter)
			if err != nil {
				panic(err)
			}
			var results []models.Appointment
			if err = cursor.All(context.TODO(), &results); err != nil {
				panic(err)
			}

			c.HTML(http.StatusOK, "AdminAppointmentList.html", gin.H{
				"title": "Manage Appointments",
				"list":  results,
			})
		}
	}
}

func EditAdminAppointment() gin.HandlerFunc {
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
			appointmentId := c.Param("appointmentId")
			objId, _ := primitive.ObjectIDFromHex(appointmentId)
			var userHairCompany models.UserHairCompany
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var appointment models.Appointment
			defer cancel()
			err := userHairCompanyCollection.FindOne(ctx, bson.M{"userId": user.ID}).Decode(&userHairCompany)
			rootUrl := "/appointment/admin"
			formTitle := "Create appointment"
			if err != nil {
				c.Redirect(http.StatusFound, "/myCompany")
			}
			appointmentCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&appointment)
			if appointment.ID.Hex() == appointmentId {
				rootUrl = "/appointment/admin/edit/" + appointmentId
				formTitle = "Update appointment"

			}

			filter := bson.D{primitive.E{Key: "hairCompanyId", Value: userHairCompany.HairCompanyId}}

			// Retrieves documents that match the query filer
			cursor, err := hairdresserCollection.Find(context.TODO(), filter)
			if err != nil {
				panic(err)
			}
			var hairdressers []models.Hairdresser
			if err = cursor.All(context.TODO(), &hairdressers); err != nil {
				panic(err)
			}

			cursor2, err := appointmentTypeCollection.Find(context.TODO(), filter)
			if err != nil {
				panic(err)
			}
			var types []models.AppointmentType
			if err = cursor2.All(context.TODO(), &types); err != nil {
				panic(err)
			}

			c.HTML(http.StatusOK, "editAppointment.html", gin.H{
				"title":         "Manage Appointment",
				"rootUrl":       rootUrl,
				"formTitle":     formTitle,
				"aTitle":        appointment.Title,
				"price":         appointment.Price,
				"date":          appointment.Date,
				"time":          appointment.Time,
				"typeId":        appointment.TypeId,
				"hairdresserId": appointment.HairdresserId,
				"types":         types,
				"hairdressers":  hairdressers,
			})
		}
	}
}

func Country() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user = &models.User{}
		session, _ := store.Get(c.Request, "session")
		val := session.Values["user"]
		var ok bool
		fmt.Println(user)
		if user, ok = val.(*models.User); !ok {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Signin / Signup",
			})
		} else {
			c.HTML(http.StatusOK, "publicCountryList.html", gin.H{
				"title": "Choice Country",
			})
		}
	}
}

func City() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user = &models.User{}
		session, _ := store.Get(c.Request, "session")
		val := session.Values["user"]
		var ok bool
		fmt.Println(user)
		if user, ok = val.(*models.User); !ok {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Signin / Signup",
			})
		} else {
			c.HTML(http.StatusOK, "publicCityList.html", gin.H{
				"title": "Choice City",
			})
		}
	}
}

func Company() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user = &models.User{}
		session, _ := store.Get(c.Request, "session")
		val := session.Values["user"]
		var ok bool
		fmt.Println(user)
		if user, ok = val.(*models.User); !ok {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Signin / Signup",
			})
		} else {
			city := c.Param("city")
			filter := bson.D{primitive.E{Key: "address.city", Value: city}}
			cursor, err := hairCompanyCollection.Find(context.TODO(), filter)
			if err != nil {
				panic(err)
			}
			var results []models.HairCompany
			if err = cursor.All(context.TODO(), &results); err != nil {
				panic(err)
			}
			c.HTML(http.StatusOK, "publicCompanyList.html", gin.H{
				"title": "Choice Company",
				"list":  results,
			})
		}
	}
}

func Type() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user = &models.User{}
		session, _ := store.Get(c.Request, "session")
		val := session.Values["user"]
		var ok bool
		fmt.Println(user)
		if user, ok = val.(*models.User); !ok {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Signin / Signup",
			})
		} else {
			companyId := c.Param("companyId")
			objId, _ := primitive.ObjectIDFromHex(companyId)
			filter := bson.D{primitive.E{Key: "hairCompanyId", Value: objId}}
			cursor, err := appointmentTypeCollection.Find(context.TODO(), filter)
			if err != nil {
				panic(err)
			}
			var results []models.AppointmentType
			if err = cursor.All(context.TODO(), &results); err != nil {
				panic(err)
			}
			c.HTML(http.StatusOK, "publicTypeList.html", gin.H{
				"title": "Choice Type",
				"list":  results,
			})
		}
	}
}

func Hairdresser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user = &models.User{}
		session, _ := store.Get(c.Request, "session")
		val := session.Values["user"]
		var ok bool
		fmt.Println(user)
		if user, ok = val.(*models.User); !ok {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Signin / Signup",
			})
		} else {
			typeId := c.Param("typeId")
			objId, _ := primitive.ObjectIDFromHex(typeId)
			filter := bson.D{primitive.E{Key: "typeId", Value: objId}}
			cursor, err := hairdresserCollection.Find(context.TODO(), filter)
			if err != nil {
				panic(err)
			}
			var results []models.Hairdresser
			if err = cursor.All(context.TODO(), &results); err != nil {
				panic(err)
			}
			c.HTML(http.StatusOK, "publicHairdresserList.html", gin.H{
				"title": "Choice Hairdresser",
				"list":  results,
			})
		}
	}
}

func Appointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user = &models.User{}
		session, _ := store.Get(c.Request, "session")
		val := session.Values["user"]
		var ok bool
		fmt.Println(user)
		if user, ok = val.(*models.User); !ok {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Signin / Signup",
			})
		} else {
			typeId := c.Param("typeId")
			hairdresserId := c.Param("hairdresserId")
			objId, _ := primitive.ObjectIDFromHex(typeId)
			objId2, _ := primitive.ObjectIDFromHex(hairdresserId)
			filter := bson.D{primitive.E{Key: "typeId", Value: objId}, primitive.E{Key: "hairdresserId", Value: objId2}}
			cursor, err := appointmentCollection.Find(context.TODO(), filter)
			if err != nil {
				panic(err)
			}
			var results []models.Appointment
			if err = cursor.All(context.TODO(), &results); err != nil {
				panic(err)
			}
			c.HTML(http.StatusOK, "publicAppointmentList.html", gin.H{
				"title": "Choice Appointment",
				"list":  results,
			})
		}
	}
}