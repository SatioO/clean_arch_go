package repository

import (
	"context"

	"github.com/satioO/togo/app/domain"
	"go.mongodb.org/mongo-driver/bson"
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
func (r *PermissionRepository) Find(ctx context.Context) ([]domain.Permission, error) {
	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{})

	permissions := []domain.Permission{}
	err = cursor.All(ctx, &permissions)

	return permissions, err
}

// Save ....
func (r *PermissionRepository) Save(ctx context.Context, permission *domain.Permission) error {
	_, err := r.db.Collection(r.collection).InsertOne(ctx, &permission)
	return err
}
