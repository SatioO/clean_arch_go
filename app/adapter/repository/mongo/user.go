package repository

import (
	"context"
	"time"

	"github.com/satioO/togo/app/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository ...
type UserRepository struct {
	db         *mongo.Database
	collection string
}

// NewUserRepository ...
func NewUserRepository(db *mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		db,
		collection,
	}
}

// FindAll ....
func (r *UserRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	pipeline := []bson.M{
		bson.M{
			"$lookup": bson.M{
				"from":         "address",
				"as":           "address",
				"localField":   "_id",
				"foreignField": "user_id",
			},
		},
	}

	cursor, err := r.db.Collection(r.collection).Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	users := []domain.User{}
	err = cursor.All(ctx, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// Save ....
func (r *UserRepository) Save(ctx context.Context, user *domain.User) error {
	result, err := r.db.Collection(r.collection).InsertOne(ctx, &domain.User{
		Name:      user.Name,
		Phone:     user.Phone,
		Email:     user.Email,
		Role:      user.Role,
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
	})

	if len(user.Address) > 0 {
		var address []interface{}

		for _, value := range user.Address {
			value.UserID, _ = result.InsertedID.(primitive.ObjectID)
			address = append(address, value)
		}

		_, err = r.db.Collection("address").InsertMany(ctx, address)
	}

	return err
}
