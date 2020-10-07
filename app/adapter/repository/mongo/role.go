package repository

import (
	"context"

	"github.com/satioO/togo/app/domain"
	"go.mongodb.org/mongo-driver/bson"
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

// Find ....
func (r *RoleRepository) Find(ctx context.Context) ([]domain.Role, error) {
	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{})

	roles := []domain.Role{}
	err = cursor.All(ctx, &roles)

	return roles, err
}

// Save ....
func (r *RoleRepository) Save(ctx context.Context, role *domain.Role) error {
	_, err := r.db.Collection(r.collection).InsertOne(ctx, &domain.Role{
		Name: role.Name,
	})

	return err
}
