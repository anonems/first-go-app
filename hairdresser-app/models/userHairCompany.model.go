package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserHairCompany struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId        primitive.ObjectID `bson:"userId,omitempty" json:"userId,omitempty" validate:"required"`
	HairCompanyId primitive.ObjectID `bson:"hairCompanyId,omitempty" json:"hairCompanyId,omitempty" validate:"required"`
	Type          string             `bson:"type,omitempty" json:"type,omitempty" validate:"required"` //type can be owner, guest
}
