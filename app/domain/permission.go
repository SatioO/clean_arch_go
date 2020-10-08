package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Permission ...
type Permission struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Type   string             `json:"type,omitempty" bson:"type,omitempty"`
	Policy primitive.ObjectID `json:"policy,omitempty" bson:"policy,omitempty"`
}
