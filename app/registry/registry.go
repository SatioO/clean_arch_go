package registry

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type registry struct {
	db     *mongo.Database
	router *mux.Router
}

// Registry ...
type Registry interface {
	RegisterUserHandler()
}

// NewRegistry ...
func NewRegistry(db *mongo.Database, router *mux.Router) Registry {
	return &registry{db, router}
}
