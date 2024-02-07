package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserHairCompany struct {
	Id            primitive.ObjectID `json:"id,omitempty"`
	UserId        primitive.ObjectID `json:"userId,omitempty" validate:"required"`
	HairCompanyId primitive.ObjectID `json:"hairCompaniyId,omitempty" validate:"required"`
	Type          string             `json:"status,omitempty" validate:"required"` //type can be owner, guest
}
