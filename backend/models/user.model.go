package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id               primitive.ObjectID   `json:"id,omitempty"`
	Email            string               `json:"email,omitempty" validate:"required"`
	Password         string               `json:"password,omitempty" validate:"required"`
	FirstName        string               `json:"firstName,omitempty" validate:"required"`
	LastName         string               `json:"lastName,omitempty" validate:"required"`
	FavHairCompanies []primitive.ObjectID `json:"favHairCompanies,omitempty"`
}
