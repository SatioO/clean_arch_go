package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Resource ...
type Resource struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Effect string             `json:"effect,omitempty" bson:"effect,omitempty"`
	Action []AccessLevel      `json:"action,omitempty" bson:"action,omitempty"`
}
