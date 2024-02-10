package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppointmentType struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name          string             `bson:"name,omitempty" json:"name,omitempty" validate:"required"`
	Description   string             `bson:"description,omitempty" json:"description,omitempty" validate:"required"`
	Duration      int32              `bson:"duration,omitempty" json:"duration,omitempty" validate:"required"`
	HairCompanyId primitive.ObjectID `bson:"hairCompanyId,omitempty" json:"hairCompanyId,omitempty" validate:"required"`
}
