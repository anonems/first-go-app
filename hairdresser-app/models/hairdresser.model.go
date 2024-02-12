package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hairdresser struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	HairCompanyId primitive.ObjectID `bson:"hairCompanyId,omitempty" json:"hairCompanyId,omitempty" validate:"required"`
	FirstName     string             `bson:"firstName,omitempty" json:"firstName,omitempty" validate:"required"`
	LastName      string             `bson:"lastName,omitempty" json:"lastName,omitempty" validate:"required"`
	TypeId        primitive.ObjectID `bson:"typeId,omitempty" json:"typeId,omitempy" validate:"required"`
}
