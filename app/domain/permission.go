package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AccessLevel ...
type AccessLevel string

const (
	// ReadAccess ...
	ReadAccess AccessLevel = "Read"
	// WriteAccess ...
	WriteAccess AccessLevel = "Write"
	// ListAccess ...
	ListAccess AccessLevel = "List"
	// FullAccess ...
	FullAccess AccessLevel = "FullAccess"
)

// Permission ...
type Permission struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	AccessLevel []AccessLevel      `json:"accessLevel" bson:"accessLevel"`
}
