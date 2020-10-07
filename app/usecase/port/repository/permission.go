package repository

import "context"

// PermissionRepository ...
type PermissionRepository interface {
	Find(context.Context) error
	Save(context.Context) error
}
