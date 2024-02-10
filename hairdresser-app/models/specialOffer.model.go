package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SpecialOffer struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" bson:"_id" json:"_id"`
	Name           string             `bson:"name,omitempty" json:"name,omitempty" validate:"required"`
	Description    string             `bson:"description,omitempty" json:"description,omitempty" validate:"required"`
	Code           string             `bson:"code,omitempty" json:"code,omitempty" validate:"required"`
	Value          float32            `bson:"value,omitempty" json:"value,omitempty" validate:"required"`
	ExpirationDate primitive.DateTime `bson:"expirationDate,omitempty" json:"expirationDate,omitempty" validate:"required"`
	SpecialOfferId primitive.ObjectID `bson:"specialOfferId,omitempty" json:"specialOfferId,omitempty"`
	HairCompanyId  primitive.ObjectID `bson:"hairCompanyId,omitempty" json:"hairCompanyId,omitempty" validate:"required"`
}
