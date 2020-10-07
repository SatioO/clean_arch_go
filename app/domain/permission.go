package domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AccessLevel ...
type AccessLevel string

const (
	// ReadAccess ...
	ReadAccess AccessLevel = "Read"
	// WriteAccess ...
	WriteAccess = "Write"
	// ListAccess ...
	ListAccess = "List"
	// FullAccess ...
	FullAccess = "FullAccess"
)

// Permission ...
type Permission struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	AccessLevel []AccessLevel      `json:"accessLevel" bson:"accessLevel"`
}

// ValidateAccessLevel ...
func ValidateAccessLevel(access []AccessLevel) error {
	var err error

	for _, level := range access {
		if level != ReadAccess && level != WriteAccess && level != ListAccess && level != FullAccess {
			err = errors.New("AccessLevel Mismatched")
		}
	}

	return err
}
