package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// PermissionRepository ...
type PermissionRepository struct {
	db         *mongo.Database
	collection string
}

// NewPermissionRepository ...
func NewPermissionRepository(db *mongo.Database, collection string) *PermissionRepository {
	return &PermissionRepository{db, collection}
}

// Find ....
func (r *PermissionRepository) Find(ctx context.Context) error {
	return nil
}

// Save ....
func (r *PermissionRepository) Save(ctx context.Context) error {
	return nil
}
