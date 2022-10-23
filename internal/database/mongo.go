package database

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	RentalsCollection = "rentals"
)

type Mongo struct {
	ConnectionString string
	DatabaseName     string

	clientInstance  *mongoDriver.Client
	defaultDatabase *mongoDriver.Database
}

func (mongo *Mongo) GetClient() *mongoDriver.Client {
	if mongo.clientInstance == nil {
		panic("mongo client is not connected.")
	}
	return mongo.clientInstance
}

func (mongo *Mongo) Collection(name string) *mongoDriver.Collection {
	return mongo.Database().Collection(name)
}

// Database returns the default database instance
func (mongo *Mongo) Database() *mongoDriver.Database {
	if mongo.defaultDatabase == nil {
		panic("Mongo database not found")
	}
	return mongo.defaultDatabase
}

func (mongo *Mongo) Connect(ctx context.Context) error {
	var err error

loop:
	for {
		select {
		case <-ctx.Done():
			return err
		default:
			err = nil
			// Set client options
			clientOptions := options.Client().ApplyURI(mongo.ConnectionString)
			// Connect to MongoDB
			client, err := mongoDriver.Connect(ctx, clientOptions)
			if err != nil {
				log.Warn().Err(err).Msg("error while connecting to mongodb")
				<-time.After(1 * time.Second)
				continue
			}
			// Check the connection
			err = client.Ping(ctx, nil)
			if err != nil {
				log.Warn().Err(err).Msg("error while connecting to mongodb")
				<-time.After(1 * time.Second)
				continue
			}

			mongo.clientInstance = client
			mongo.defaultDatabase = client.Database(mongo.DatabaseName)

			break loop
		}
	}

	return err
}
