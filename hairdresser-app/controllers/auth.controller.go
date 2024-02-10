package controllers

import (
	"context"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"hairdresser-app/models"
	"hairdresser-app/utils"
	"net/http"
	"time"
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
			Username: c.Request.PostFormValue("username"),
			Password: c.Request.PostFormValue("password"),
		}

		//validate the request body
		if err := c.BindQuery(&login); err != nil {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"title":        "Signin / Signup",
				"errorMessage": err.Error(),
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&login); validationErr != nil {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"title":        "Signin / Signup",
				"errorMessage": validationErr.Error(),
			})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": login.Username}).Decode(&user)
		if err != nil {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"title":        "Signin / Signup",
				"errorMessage": err.Error(),
			})
			return
		}

		if utils.CheckPasswordHash(login.Password, user.Password) {
			//configure session
			session, _ := store.Get(c.Request, "session")
			session.Values["user"] = user
			session.Save(c.Request, c.Writer)

			c.Redirect(http.StatusFound, "/")
		} else {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"title":        "Signin / Signup",
				"errorMessage": "Invalid password / username",
			})
			return
		}

	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := store.Get(c.Request, "session")
		delete(session.Values, "user")
		session.Save(c.Request, c.Writer)
		c.Redirect(http.StatusFound, "/")
	}
}
