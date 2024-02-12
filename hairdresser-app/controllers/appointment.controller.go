package controllers

import (
	"context"
	"fmt"
	"hairdresser-app/constants"
	"hairdresser-app/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//var validate = validator.New()

func CreateAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		price, _ := strconv.Atoi(c.Request.PostFormValue("price"))
		typeId, _ := primitive.ObjectIDFromHex(c.Request.PostFormValue("typeId"))
		hairdresserId, _ := primitive.ObjectIDFromHex(c.Request.PostFormValue("hairdresserId"))
		appointment := models.Appointment{
			Title: c.Request.PostFormValue("title"),
			Time: c.Request.PostFormValue("time"),
			Date: c.Request.PostFormValue("date"),
			Price: float32(price),
			TypeId: typeId,
			HairdresserId: hairdresserId,
			Status: constants.AVAILABLE,
		}

		//validate the request body
		if err := c.BindQuery(&appointment); err != nil {
			c.HTML(http.StatusBadRequest, "editAppointment.html", gin.H{
				"formTitle":    "Create Appointment",
				"title":        "Manage Appointments",
				"errorMessage": err.Error(),
				"rootUrl":      "/appointment/admin",
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

		newAppointment := models.Appointment{
			Title: appointment.Title,
			Time: appointment.Time,
			Date: appointment.Date,
			Price: appointment.Price,
			TypeId: appointment.TypeId,
			HairdresserId: appointment.HairdresserId,
			Status: appointment.Status,
			HairCompanyId: userHairCompany.HairCompanyId,
		}

		result, err := appointmentCollection.InsertOne(ctx, newAppointment)
		if err != nil {
			c.HTML(http.StatusBadRequest, "editAppointment.html", gin.H{
				"formTitle":    "Create Appointment",
				"title":        "Manage Appointment",
				"errorMessage": err.Error(),
				"rootUrl":      "/appointment/admin",
			})
			return
		}
		fmt.Println(result)
		//oid, _ := result.InsertedID.(primitive.ObjectID)

		c.Redirect(http.StatusFound, "/appointment/admin")

	}
}

func EditAppointment() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		appointmentId := c.Param("appointmentId")
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(appointmentId)

		price, _ := strconv.Atoi(c.Request.PostFormValue("price"))
		typeId, _ := primitive.ObjectIDFromHex(c.Request.PostFormValue("typeId"))
		hairdresserId, _ := primitive.ObjectIDFromHex(c.Request.PostFormValue("hairdresserId"))
		appointment := models.Appointment{
			Title: c.Request.PostFormValue("title"),
			Time: c.Request.PostFormValue("time"),
			Date: c.Request.PostFormValue("date"),
			Price: float32(price),
			TypeId: typeId,
			HairdresserId: hairdresserId,
			Status: constants.AVAILABLE,
		}

		//validate the request body
		if err := c.BindQuery(&appointment); err != nil {
			c.HTML(http.StatusBadRequest, "editAppointment.html", gin.H{
				"formTitle":    "Create Appointment",
				"title":        "Manage Appointment",
				"errorMessage": err.Error(),
			})
			return
		}

		update := bson.M{
			"title": appointment.Title,
			"time": appointment.Time,
			"date": appointment.Date,
			"price": appointment.Price,
			"typeId": appointment.TypeId,
			"hairdresserId": appointment.HairdresserId,
			"status": appointment.Status,
		}

		result, err := appointmentCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.HTML(http.StatusBadRequest, "editAppointment.html", gin.H{
				"formTitle":    "Update Appointment",
				"title":        "Manage Appointment",
				"errorMessage": err.Error(),
				"rootUrl":      "/appointment/edit/" + appointmentId,
			})
			return
		}

		//get updated details
		var updatedAppointment models.Appointment
		if result.MatchedCount == 1 {
			err := appointmentCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedAppointment)
			if err != nil {
				c.HTML(http.StatusBadRequest, "editAppointment.html", gin.H{
					"formTitle":    "Update Type",
					"title":        "Manage Appointment Types",
					"errorMessage": err.Error(),
					"rootUrl":      "/appointment/edit/admin/" + appointmentId,
				})
				return
			}
		}

		c.Redirect(http.StatusFound, "/appointment/admin")

	}
}

func DeleteAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		appointmentId := c.Param("appointmentId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(appointmentId)

		appointmentCollection.DeleteOne(ctx, bson.M{"_id": objId})
		
		c.Redirect(http.StatusFound, "/appointment/admin")
	}
}

func TakeAppointment() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		appointmentId := c.Param("appointmentId")
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(appointmentId)
		var user = &models.User{}
		session, _ := store.Get(c.Request, "session")
		val := session.Values["user"]
		fmt.Println(user)
		user, _= val.(*models.User)
		appointment := models.Appointment{
			UserId: user.ID,
			Status: constants.RESERVED,
		}

		update := bson.M{
			"userId": appointment.UserId,
			"status": appointment.Status,
		}

		appointmentCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		c.Redirect(http.StatusFound, "/")

	}
}