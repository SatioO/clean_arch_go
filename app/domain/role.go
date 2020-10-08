package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Role entity
type Role struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Policy []Policy           `json:"policy,omitempty" bson:"policy,omitempty"`
}
