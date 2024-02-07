package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Appointment struct {
	Id             primitive.ObjectID `json:"id,omitempty"`
	HairCompanyId  primitive.ObjectID `json:"hairCompanyId,omitempty" validate:"required"`
	Status         string             `json:"status,omitempty" validate:"required"` //status can be available, reserved, expired or cancelled,
	Time           string             `json:"time,omitempty" validate:"required"`
	Date           primitive.DateTime `json:"date,omitempty" validate:"required"`
	UserId         primitive.ObjectID `json:"userId,omitempy"`
	Price          float32            `json:"price,omitempy" validate:"required"`
	Title          string             `json:"title,omitempy" validate:"required"`
	TypeId         primitive.ObjectID `json:"typeId,omitempy" validate:"required"`
	SpecialOfferId primitive.ObjectID `json:"specialOfferId,omitempy"`
}
