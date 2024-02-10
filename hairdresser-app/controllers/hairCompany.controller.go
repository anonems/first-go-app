package controllers

import (
	"hairdresser-app/constants"
	"hairdresser-app/models"
	"hairdresser-app/responses"

	//"hairdresser-app/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	//"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//var validate = validator.New()

func CreateHairCompany() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		address := models.Address{
			Line1:      c.Request.PostFormValue("line1"),
			Line2:      c.Request.PostFormValue("line2"),
			PostalCode: c.Request.PostFormValue("postalCode"),
			City:       c.Request.PostFormValue("city"),
			Country:    c.Request.PostFormValue("country"),
		}

		hairCompany := models.HairCompany{
			Name:    c.Request.PostFormValue("name"),
			SIREN:   c.Request.PostFormValue("siren"),
			Address: address,
		}

		//validate the request body
		if err := c.BindQuery(&hairCompany); err != nil {
			c.HTML(http.StatusBadRequest, "myCompany.html", gin.H{
				"errorMessage": err.Error(),
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&hairCompany); validationErr != nil {
			c.HTML(http.StatusBadRequest, "myCompany.html", gin.H{
				"title":        "My Company",
				"errorMessage": validationErr.Error(),
			})
			return
		}

		newHairCompany := models.HairCompany{
			Name:    hairCompany.Name,
			SIREN:   hairCompany.SIREN,
			Status:  constants.OPPENED,
			Address: hairCompany.Address,
		}

		result, err := hairCompanyCollection.InsertOne(ctx, newHairCompany)
		if err != nil {
			c.HTML(http.StatusBadRequest, "myCompany.html", gin.H{
				"title":        "My Company",
				"errorMessage": err.Error(),
			})
			return
		}

		oid, _ := result.InsertedID.(primitive.ObjectID)


		var user = &models.User{}
		session, _ := store.Get(c.Request, "session")
		val := session.Values["user"]
		user, _ = val.(*models.User)
		newUserHairCompany := models.UserHairCompany{
			HairCompanyId: oid,
			UserId:        user.ID,
			Type:          constants.OWNER,
		}
		_, err2 := userHairCompanyCollection.InsertOne(ctx, newUserHairCompany)
		if err2 != nil {
			c.HTML(http.StatusBadRequest, "myCompany.html", gin.H{
				"title":        "My Company",
				"errorMessage": err2.Error(),
			})
			return
		}

		c.HTML(http.StatusOK, "myCompany.html", gin.H{
			"title":          "My Company",
			"rootUrl":        "/hairCompany/" + oid.Hex(),
			"formTitle":      "Update Company",
			"name":           hairCompany.Name,
			"siren":          hairCompany.SIREN,
			"line1":          hairCompany.Address.Line1,
			"line2":          hairCompany.Address.Line2,
			"postalCode":     hairCompany.Address.PostalCode,
			"city":           hairCompany.Address.City,
			"country":        hairCompany.Address.Country,
			"successMessage": "Your company has been created!",
		})

	}
}

func GetHairCompany() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

func EditHairCompany() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		companyId := c.Param("companyId")
		defer cancel()

		address := models.Address{
			Line1:      c.Request.PostFormValue("line1"),
			Line2:      c.Request.PostFormValue("line2"),
			PostalCode: c.Request.PostFormValue("postalCode"),
			City:       c.Request.PostFormValue("city"),
			Country:    c.Request.PostFormValue("country"),
		}

		hairCompany := models.HairCompany{
			Name:    c.Request.PostFormValue("name"),
			SIREN:   c.Request.PostFormValue("siren"),
			Address: address,
		}

		//validate the request body
		if err := c.BindQuery(&hairCompany); err != nil {
			c.HTML(http.StatusBadRequest, "myCompany.html", gin.H{
				"errorMessage": err.Error(),
			})
			return
		}

		update := bson.M{
			"name":    hairCompany.Name,
			"siren":   hairCompany.SIREN,
			"address": hairCompany.Address,
		}

		_, err := userCollection.UpdateOne(ctx, bson.M{"id": companyId}, bson.M{"$set": update})
		if err != nil {
			c.HTML(http.StatusBadRequest, "myCompany.html", gin.H{
				"title":        "My Company",
				"errorMessage": err.Error(),
			})
			return
		}

		c.HTML(http.StatusOK, "myCompany.html", gin.H{
			"title":          "My Company",
			"rootUrl":        "/hairCompany/" + hairCompany.ID.Hex(),
			"formTitle":      "Update Company",
			"name":           hairCompany.Name,
			"siren":          hairCompany.SIREN,
			"line1":          hairCompany.Address.Line1,
			"line2":          hairCompany.Address.Line2,
			"postalCode":     hairCompany.Address.PostalCode,
			"city":           hairCompany.Address.City,
			"country":        hairCompany.Address.Country,
			"successMessage": "Your company has been updated!",
		})

	}
}

func DeleteHairCompany() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.DefaultResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()},
				},
			)
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.DefaultResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
		)
	}
}

func GetAllHairCompanies() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			users = append(users, singleUser)
		}

		c.JSON(http.StatusOK,
			responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
		)
	}
}
