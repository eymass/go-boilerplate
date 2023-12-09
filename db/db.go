package db

import (
	"context"
	"github.com/heroku/go-getting-started/common/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const (
	dbNameVar = "DATABASE_NAME"
	dbUriVar  = "MONGODB_URI"
)

func GetDatabase(ctx context.Context) (*mongo.Database, func()) {
	dbUri := os.Getenv(dbUriVar)
	if dbUri == "" {
		logger.ErrorLogger.Printf("could not find database uri")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUri))

	if err != nil {
		logger.ErrorLogger.Fatal(err)
	}

	// do connect
	err = client.Connect(ctx)
	if err != nil {
		logger.ErrorLogger.Fatal(err)
	}

	// cleanup
	cleanup := func() {
		client.Disconnect(ctx)
	}

	// get database
	dbName := os.Getenv(dbNameVar)
	if dbName == "" {
		logger.ErrorLogger.Printf("could not find database name")
	}
	return client.Database(dbName), cleanup
}
