package repository

import (
	"context"

	"github.com/satioO/togo/app/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PolicyRepository ...
type PolicyRepository struct {
	db         *mongo.Database
	collection string
}

// NewPolicyRepository ...
func NewPolicyRepository(db *mongo.Database, collection string) *PolicyRepository {
	return &PolicyRepository{db, collection}
}

// Find ....
func (r *PolicyRepository) Find(ctx context.Context) ([]domain.Policy, error) {
	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{})

	policies := []domain.Policy{}
	err = cursor.All(ctx, &policies)

	return policies, err
}

// Save ....
func (r *PolicyRepository) Save(ctx context.Context, Policy *domain.Policy) error {
	_, err := r.db.Collection(r.collection).InsertOne(ctx, &Policy)
	return err
}
