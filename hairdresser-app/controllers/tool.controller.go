package controllers

import (
	"hairdresser-app/configs"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

//mongo collection
var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var userHairCompanyCollection *mongo.Collection = configs.GetCollection(configs.DB, "userHairCompanies")
var hairCompanyCollection *mongo.Collection = configs.GetCollection(configs.DB, "hairCompanies")

//validator
var validate = validator.New()