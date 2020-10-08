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

// Policy ...
type Policy struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Type     string             `json:"type,omitempty" bson:"type,omitempty"`
	Effect   string             `json:"effect,omitempty" bson:"effect,omitempty"`
	Action   []AccessLevel      `json:"action,omitempty" bson:"action,omitempty"`
	Resource []string           `json:"resource,omitempty" bson:"resource,omitempty"`
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
