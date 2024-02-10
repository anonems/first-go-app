package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hairdresser-app/models"
	"hairdresser-app/responses"
	"hairdresser-app/utils"
	"net/http"
	"strconv"
	"time"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		age, _ := strconv.Atoi(c.Request.PostFormValue("age"))
		user := models.User{
			Email:     c.Request.PostFormValue("email"),
			Password:  c.Request.PostFormValue("password"),
			FirstName: c.Request.PostFormValue("firstName"),
			LastName:  c.Request.PostFormValue("lastName"),
			Age:       age,
			Gender:    c.Request.PostFormValue("gender"),
		}
		// filter := bson.D{{"email", user.Email}}
		// var checkEmail []models.User
		// cursor, _ := userCollection.Find(context.TODO(), filter)
		// cursor.All(context.TODO(), &checkEmail)
		// if len(checkEmail) > 0 {
		// 	c.JSON(http.StatusNotAcceptable, responses.DefaultResponse{Status: http.StatusInternalServerError})
		// 	return
		// }

		//validate the request body
		if err := c.BindQuery(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.DefaultResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.DefaultResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		pwdHash, _ := utils.HashPassword(user.Password)

		newUser := models.User{
			Email:     user.Email,
			Password:  pwdHash,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Age:       user.Age,
			Gender:    user.Gender,
		}

		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.DefaultResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		fmt.Println(userId)
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}

func EditUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()

		age, _ := strconv.Atoi(c.Request.PostFormValue("age"))
		user := models.User{
			Email:     c.Request.PostFormValue("email"),
			Password:  c.Request.PostFormValue("password"),
			FirstName: c.Request.PostFormValue("firstName"),
			LastName:  c.Request.PostFormValue("lastName"),
			Age:       age,
			Gender:    c.Request.PostFormValue("gender"),
		} 
		objId, _ := primitive.ObjectIDFromHex(userId)
		// filter := bson.D{{"email", user.Email}}
		// var checkEmail []models.User
		// cursor, _ := userCollection.Find(context.TODO(), filter)
		// cursor.All(context.TODO(), &checkEmail)
		// if len(checkEmail) > 0 {
		// 	c.JSON(http.StatusNotAcceptable, responses.DefaultResponse{Status: http.StatusInternalServerError})
		// 	return
		// }
		// fmt.Println(objId)

		//validate the request body
		if err := c.BindQuery(&user); err != nil {
			c.HTML(http.StatusBadRequest, "myUser.html", gin.H{
				"errorMessage": err.Error(),
			})
			return
		}

		update := bson.M{
			"email":     user.Email,
			//"password":  user.Password,
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"age":       user.Age,
			"gender":    user.Gender,
		}
		result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.HTML(http.StatusBadRequest, "myUser.html", gin.H{
				"errorMessage": err.Error(),
			})
			return
		}

		//get updated user details
		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
			if err != nil {
				c.HTML(http.StatusBadRequest, "myUser.html", gin.H{
					"errorMessage": err.Error(),
				})
				return
			}
		}

		c.HTML(http.StatusOK, "myProfile.html", gin.H{
			"title":          "My Profile",
			"userId":         updatedUser.ID,
			"firstName":      updatedUser.FirstName,
			"lastName":       updatedUser.LastName,
			"age":            updatedUser.Age,
			"email":          updatedUser.Email,
			"gender":         updatedUser.Gender,
			"successMessage": "Your profile has been updated!",
		})

	}
}
