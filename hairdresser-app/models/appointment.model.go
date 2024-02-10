package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Appointment struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	HairCompanyId  primitive.ObjectID `bson:"hairCompanyId,omitempty" json:"hairCompanyId,omitempty" validate:"required"`
	Status         string             `bson:"status,omitempty" json:"status,omitempty" validate:"required"` //status can be available, reserved, expired or cancelled,
	Time           string             `bson:"time,omitempty" json:"time,omitempty" validate:"required"`
	Date           primitive.DateTime `bson:"date,omitempty" json:"date,omitempty" validate:"required"`
	UserId         primitive.ObjectID `bson:"userId,omitempty" json:"userId,omitempy"`
	Price          float32            `bson:"price,omitempty" json:"price,omitempy" validate:"required"`
	Title          string             `bson:"title,omitempty" json:"title,omitempy" validate:"required"`
	TypeId         primitive.ObjectID `bson:"typeId,omitempty" json:"typeId,omitempy" validate:"required"`
	SpecialOfferId primitive.ObjectID `bson:"specialOfferId,omitempty" json:"specialOfferId,omitempy"`
}
