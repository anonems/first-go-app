package controllers

import (
	"context"
	"hairdresser-app/models"
	"hairdresser-app/responses"
	"hairdresser-app/utils"
	"net/http"
	"time"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)
	
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		login := models.Login{
			Username : c.Request.PostFormValue("username"),
			Password : c.Request.PostFormValue("password"),
		}


		//validate the request body
		if err := c.BindQuery(&login); err != nil {
			c.JSON(http.StatusBadRequest, responses.DefaultResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&login); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.DefaultResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": login.Username}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}


		if !utils.CheckPasswordHash(login.Password, user.Password) {
			c.JSON(http.StatusBadRequest, responses.DefaultResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "invalid credential!"}})
			return
		}

		//configure session
		session := sessions.Default(c)
		session.Set("username", login.Username)
		session.Save()

		c.JSON(http.StatusOK, responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})

	}
}


func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign out successfully",
	})
}