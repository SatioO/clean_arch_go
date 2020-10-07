package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Address entity
type Address struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID  primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
	Street  string             `json:"street,omitempty" bson:"street,omitempty"`
	City    string             `json:"city,omitempty" bson:"city,omitempty"`
	State   string             `json:"state,omitempty" bson:"state,omitempty"`
	Pincode int64              `json:"pincode,omitempty" bson:"pincode,omitempty"`
}
