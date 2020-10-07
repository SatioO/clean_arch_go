package repository

import "context"

// RoleRepository ...
type RoleRepository interface {
	Save(context.Context) error
}
