package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const dbName = "graphql"

var dbURL string

func init() {
	dbURL = os.Getenv("MONGODB_URL")
	if dbURL == "" {
		log.Fatal("missing MONGODB_URL")
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	if err != nil {
		log.Fatalf("invalid dburl '%s': %v", dbURL, err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("unable to ping db '%s': %v", dbURL, err)
	}

	fmt.Println("Connected to mongodb at", dbURL)
}

func getDB() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	if err != nil {
		return nil, fmt.Errorf("could not connect to mongodb '%s': %v", dbURL, err)
	}

	return client.Database(dbName), nil
}
