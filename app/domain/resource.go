package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// EffectType ...
type EffectType string

const (
	// ALLOW ...
	ALLOW EffectType = "Allow"

	// DENY ...
	DENY = "Deny"
)

// Resource ...
type Resource struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Effect EffectType         `json:"effect,omitempty" bson:"effect,omitempty"`
}
