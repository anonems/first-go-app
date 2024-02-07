package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SpecialOffer struct {
	Id             primitive.ObjectID `json:"id,omitempty"`
	Name           string             `json:"name,omitempty" validate:"required"`
	Description    string             `json:"description,omitempty" validate:"required"`
	Code           string             `json:"code,omitempty" validate:"required"`
	Value          float32            `json:"value,omitempty" validate:"required"`
	ExpirationDate primitive.DateTime `json:"expirationDate,omitempty" validate:"required"`
	SpecialOfferId primitive.ObjectID `json:"specialOfferId,omitempty"`
	HairCompanyId  primitive.ObjectID `json:"hairCompanyId,omitempty" validate:"required"`
}
