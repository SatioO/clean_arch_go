package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// RoleRepository ...
type RoleRepository struct {
	db         *mongo.Database
	collection string
}

// NewRoleRepository ...
func NewRoleRepository(db *mongo.Database, collection string) *RoleRepository {
	return &RoleRepository{db, collection}
}

// Save ....
func (r *RoleRepository) Save(ctx context.Context) error {
	return nil
}
