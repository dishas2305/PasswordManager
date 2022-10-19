package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CustomersCollection = "customers"

type CustomersModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Phone     string             `bson:"phone,omitempty" json:"phone,omitempty"`
	MPin      string             `bson:"mpin,omitempty" json:"mpin,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	OTP       string             `bson:"otp" json:"otp"`
}
