package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HairCompany struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name    string             `bson:"name,omitempty" json:"name,omitempty" validate:"required"`
	SIREN   string             `bson:"siren,omitempty" json:"siren,omitempty" validate:"required"`
	Address Address            `bson:"address,omitempty" json:"address,omitempty" validate:"required"`
	Status  string             `bson:"status,omitempty" json:"status,omitempty" Validate:"required"` //status can be oppened or closed
}

type Address struct {
	Line1      string `bson:"line1,omitempty" json:"line1,omitempty" validate:"required"`
	Line2      string `bson:"line2,omitempty" json:"line2,omitempty"`
	PostalCode string `bson:"postalCode,omitempty" json:"postalCode,omitempty" validate:"required"`
	City       string `bson:"city,omitempty" json:"city,omitempty" validate:"required"`
	Country    string `bson:"country,omitempty" json:"country,omitempty" validate:"required"`
}
