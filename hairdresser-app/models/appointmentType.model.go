package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppointmentType struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          string             `json:"name,omitempty" validate:"required"`
	Description   string             `json:"description,omitempty" validate:"required"`
	Duration      string             `json:"duration,omitempty" validate:"required"`
	Code          string             `json:"code,omitempty" validate:"required"`
	HairCompanyId primitive.ObjectID `json:"hairCompanyId,omitempty" validate:"required"`
}
