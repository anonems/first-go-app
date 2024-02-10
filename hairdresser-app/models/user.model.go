package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID               primitive.ObjectID   `bson:"_id,omitempty" json:"_id"`
	Email            string               `bson:"email,omitempty" json:"email,omitempty" validate:"required"`
	Password         string               `bson:"password,omitempty" json:"password,omitempty" validate:"required"`
	FirstName        string               `bson:"firstName,omitempty" json:"firstName,omitempty" validate:"required"`
	LastName         string               `bson:"lastName,omitempty" json:"lastName,omitempty" validate:"required"`
	FavHairCompanies []primitive.ObjectID `bson:"favHairCompanies,omitempty" json:"favHairCompanies,omitempty"`
	Age              int                  `bson:"age,omitempty" json:"age,omitempty" validate:"required"`
	Gender           string               `bson:"gender,omitempty" json:"gender,omitempty" validate:"required"`
}
