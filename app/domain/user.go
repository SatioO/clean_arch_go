package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User entity
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name,omitempty"`
	Phone     int64              `json:"phone" bson:"phone,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	Address   []Address          `json:"address" bson:"address,omitempty"`
	CreatedOn time.Time          `json:"createdOn" bson:"createdOn,omitempty"`
	UpdatedOn time.Time          `json:"updatedOn" bson:"updatedOn,omitempty"`
}

// Address entity
type Address struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID  primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Street  string             `json:"street" bson:"street,omitempty"`
	City    string             `json:"city" bson:"city,omitempty"`
	State   string             `json:"state" bson:"state,omitempty"`
	Pincode int64              `json:"pincode" bson:"pincode,omitempty"`
}

// Role entity
type Role struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty"`
	Type string             `json:"type" bson:"type,omitempty"`
}
