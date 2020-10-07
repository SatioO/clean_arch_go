package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/satioO/togo/app/registry"
)

func main() {
	db := ConnectDB("togo")
	r := mux.NewRouter().StrictSlash(true)

	_registry := registry.NewRegistry(db, r)
	_registry.RegisterUserHandler()
	_registry.RegisterRoleHandler()

	fmt.Println("App started listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

// ConnectDB estabilishes the connection with the database
func ConnectDB(dbName string) *mongo.Database {
	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)

	err = client.Connect(ctx)
	//To close the connection at the end
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	return client.Database(dbName)
}
