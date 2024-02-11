package controllers

import (
	"context"
	"hairdresser-app/models"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//var validate = validator.New()

func CreateHairdresser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		hairdresser := models.Hairdresser{
			FirstName: c.Request.PostFormValue("firstName"),
			LastName:  c.Request.PostFormValue("lastName"),
		}

		//validate the request body
		if err := c.BindQuery(&hairdresser); err != nil {
			c.HTML(http.StatusBadRequest, "editHairdresser.html", gin.H{
				"formTitle":    "Create Hairdresser",
				"title":        "Manage Hairdresser",
				"errorMessage": err.Error(),
				"rootUrl":      "/hairdresser",
			})
			return
		}

		var userHairCompany models.UserHairCompany
		var user = &models.User{}
		session, _ := store.Get(c.Request, "session")
		val := session.Values["user"]
		user, _ = val.(*models.User)
		err := userHairCompanyCollection.FindOne(ctx, bson.M{"userId": user.ID}).Decode(&userHairCompany)
		if err != nil {
			c.Redirect(http.StatusFound, "/myCompany")

		}

		newAppointmentType := models.Hairdresser{
			FirstName:     hairdresser.FirstName,
			LastName:      hairdresser.LastName,
			HairCompanyId: userHairCompany.HairCompanyId,
		}

		result, err := hairdresserCollection.InsertOne(ctx, newAppointmentType)
		if err != nil {
			c.HTML(http.StatusBadRequest, "editHairdresser.html", gin.H{
				"formTitle":    "Create Hairdresser",
				"title":        "Manage Hairdresser",
				"errorMessage": err.Error(),
				"rootUrl":      "/hairdresser",
			})
			return
		}

		oid, _ := result.InsertedID.(primitive.ObjectID)

		c.HTML(http.StatusOK, "editHairdresser.html", gin.H{
			"title":          "Manage Hairdresser",
			"rootUrl":        "/hairdresser/edit/" + oid.Hex(),
			"formTitle":      "Update Hairdresser",
			"firstName":      hairdresser.FirstName,
			"lastName":       hairdresser.LastName,
			"successMessage": "Hairdresser has been created!",
		})

	}
}

func EditHairdresser() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		hairdresserId := c.Param("hairdresserId")
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(hairdresserId)

		hairdresser := models.Hairdresser{
			FirstName: c.Request.PostFormValue("firstName"),
			LastName:  c.Request.PostFormValue("lastName"),
		}

		//validate the request body
		if err := c.BindQuery(&hairdresser); err != nil {
			c.HTML(http.StatusBadRequest, "editType.html", gin.H{
				"formTitle":    "Create Type",
				"title":        "Manage Appointment Types",
				"errorMessage": err.Error(),
			})
			return
		}

		update := bson.M{
			"firstName": hairdresser.FirstName,
			"lastName":  hairdresser.LastName,
		}

		result, err := hairdresserCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.HTML(http.StatusBadRequest, "editHairdresser.html", gin.H{
				"formTitle":    "Update Hairdresser",
				"title":        "Manage Hairdresser",
				"errorMessage": err.Error(),
				"rootUrl":      "/hairdresser/edit/" + hairdresserId,
			})
			return
		}

		//get updated details
		var updatedHairdresser models.Hairdresser
		if result.MatchedCount == 1 {
			err := hairdresserCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedHairdresser)
			if err != nil {
				c.HTML(http.StatusBadRequest, "editHairdresser.html", gin.H{
					"formTitle":    "Update Hairdresser",
					"title":        "Manage Hairdresser",
					"errorMessage": err.Error(),
					"rootUrl":      "/hairdresser/edit/" + hairdresserId,
				})
				return
			}
		}

		c.HTML(http.StatusOK, "editHairdresser.html", gin.H{
			"title":          "Manage Appointment Types",
			"rootUrl":        "/hairdresser/edit/" + hairdresserId,
			"formTitle":      "Update Hairdresser",
			"firstName":      updatedHairdresser.FirstName,
			"lastName":       updatedHairdresser.LastName,
			"successMessage": "Hairdresser has been updated!",
		})

	}
}

func DeleteHairdresser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		hairdresserId := c.Param("hairdresserId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(hairdresserId)

		hairdresserCollection.DeleteOne(ctx, bson.M{"_id": objId})

		filter := bson.D{primitive.E{Key: "hairdresserId", Value: hairdresserId}}
		_, err := appointmentCollection.DeleteMany(context.TODO(), filter)
		if err != nil {
			panic(err)
		}
		c.Redirect(http.StatusFound, "/hairdresser")
	}
}
