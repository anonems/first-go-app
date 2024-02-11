package controllers

import (
	"context"
	"hairdresser-app/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//var validate = validator.New()

func CreateAppointmentType() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		duration, _ := strconv.Atoi(c.Request.PostFormValue("duration"))
		appointmentType := models.AppointmentType{
			Name:        c.Request.PostFormValue("name"),
			Description: c.Request.PostFormValue("description"),
			Duration:    int32(duration),
		}

		//validate the request body
		if err := c.BindQuery(&appointmentType); err != nil {
			c.HTML(http.StatusBadRequest, "editType.html", gin.H{
				"formTitle":    "Create Type",
				"title":        "Manage Appointment Types",
				"errorMessage": err.Error(),
				"rootUrl":      "/appointmentType",
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

		newAppointmentType := models.AppointmentType{
			Name:          appointmentType.Name,
			Description:   appointmentType.Description,
			Duration:      appointmentType.Duration,
			HairCompanyId: userHairCompany.HairCompanyId,
		}

		result, err := appointmentTypeCollection.InsertOne(ctx, newAppointmentType)
		if err != nil {
			c.HTML(http.StatusBadRequest, "editType.html", gin.H{
				"formTitle":    "Create Type",
				"title":        "Manage Appointment Types",
				"errorMessage": err.Error(),
				"rootUrl":      "/appointmentType",
			})
			return
		}

		oid, _ := result.InsertedID.(primitive.ObjectID)

		c.HTML(http.StatusOK, "editType.html", gin.H{
			"title":          "Manage Appointment Types",
			"rootUrl":        "/appointmentType/edit/" + oid.Hex(),
			"formTitle":      "Update Type",
			"name":           appointmentType.Name,
			"description":    appointmentType.Description,
			"duration":       appointmentType.Duration,
			"successMessage": "Appointmnet type has been created!",
		})

	}
}

func EditAppointmentType() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		typeId := c.Param("typeId")
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(typeId)

		duration, _ := strconv.Atoi(c.Request.PostFormValue("duration"))
		appointmentType := models.AppointmentType{
			Name:        c.Request.PostFormValue("name"),
			Description: c.Request.PostFormValue("description"),
			Duration:    int32(duration),
		}

		//validate the request body
		if err := c.BindQuery(&appointmentType); err != nil {
			c.HTML(http.StatusBadRequest, "editType.html", gin.H{
				"formTitle":    "Create Type",
				"title":        "Manage Appointment Types",
				"errorMessage": err.Error(),
			})
			return
		}

		update := bson.M{
			"name":        appointmentType.Name,
			"description": appointmentType.Description,
			"duration":    appointmentType.Duration,
		}

		result, err := appointmentTypeCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.HTML(http.StatusBadRequest, "editType.html", gin.H{
				"formTitle":    "Update Type",
				"title":        "Manage Appointment Types",
				"errorMessage": err.Error(),
				"rootUrl":      "/appointmentType/edit/" + typeId,
			})
			return
		}

		//get updated details
		var updatedAppointmentType models.AppointmentType
		if result.MatchedCount == 1 {
			err := appointmentTypeCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedAppointmentType)
			if err != nil {
				c.HTML(http.StatusBadRequest, "editType.html", gin.H{
					"formTitle":    "Update Type",
					"title":        "Manage Appointment Types",
					"errorMessage": err.Error(),
					"rootUrl":      "/appointmentType/edit/" + typeId,
				})
				return
			}
		}

		c.HTML(http.StatusOK, "editType.html", gin.H{
			"title":          "Manage Appointment Types",
			"rootUrl":        "/appointmentType/edit/" + typeId,
			"formTitle":      "Update Type",
			"name":           updatedAppointmentType.Name,
			"description":    updatedAppointmentType.Description,
			"duration":       updatedAppointmentType.Duration,
			"successMessage": "Appointment type has been updated!",
		})

	}
}

func DeleteAppointmentType() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		typeId := c.Param("typeId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(typeId)

		appointmentTypeCollection.DeleteOne(ctx, bson.M{"_id": objId})
		
		c.Redirect(http.StatusFound, "/appointmentType")
	}
}