package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Role entity
type Role struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
	Type string             `json:"type,omitempty" bson:"type,omitempty"`
}
