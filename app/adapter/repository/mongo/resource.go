package repository

import (
	"context"

	"github.com/satioO/togo/app/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

// ResourceRepository ...
type ResourceRepository struct {
	db         *mongo.Database
	collection string
}

// NewResourceRepository ...
func NewResourceRepository(db *mongo.Database, collection string) *ResourceRepository {
	return &ResourceRepository{db, collection}
}

// Find ....
func (r *ResourceRepository) Find(ctx context.Context) ([]domain.Resource, error) {
	return nil, nil
}

// Save ....
func (r *ResourceRepository) Save(ctx context.Context, role *domain.Resource) error {
	return nil
}
