package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User entity
type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Phone     int64              `json:"phone,omitempty" bson:"phone,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Address   []Address          `json:"address,omitempty" bson:"address,omitempty"`
	CreatedOn time.Time          `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
	UpdatedOn time.Time          `json:"updatedOn,omitempty" bson:"updatedOn,omitempty"`
}
