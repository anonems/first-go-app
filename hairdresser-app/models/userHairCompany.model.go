package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserHairCompany struct {
	ID            primitive.ObjectID `bson:"_id"`
	UserId        primitive.ObjectID `json:"userId,omitempty" validate:"required"`
	HairCompanyId primitive.ObjectID `json:"hairCompaniyId,omitempty" validate:"required"`
	Type          string             `json:"type,omitempty" validate:"required"` //type can be owner, guest
}
