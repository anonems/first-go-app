package controllers

import (
	"context"
	"encoding/gob"
	"fmt"
	"hairdresser-app/models"
	"hairdresser-app/responses"
	"hairdresser-app/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
)

var store = sessions.NewCookieStore([]byte("super"))

func init() {
	store.Options.HttpOnly = true
	store.Options.Secure = true
	gob.Register(&models.User{})
}

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
		session, _ := store.Get(c.Request, "session")
		session.Values["user"] = user
		session.Save(c.Request, c.Writer)

		c.Redirect(http.StatusFound, "/")


		//c.JSON(http.StatusOK, responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})

	}
}


// func Logout(c *gin.Context) {
// 	session := sessions.Default(c)
// 	session.Clear()
// 	session.Save()
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "User Sign out successfully",
// 	})
// }

func auth(c *gin.Context) {
	fmt.Println("auth func runnig")
	session, _ := store.Get(c.Request, "session")
	fmt.Println("session:", session)
	_, ok := session.Values["user"]
	if !ok {
		c.HTML(http.StatusBadRequest, "index.html", nil)
		c.Abort()
		return
	}
	fmt.Println("func done")
	c.Next()
}