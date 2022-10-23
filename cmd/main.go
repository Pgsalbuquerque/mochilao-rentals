package main

import (
	"context"
	"log"
	"mochilao-rentals/internal/config"
	"mochilao-rentals/internal/database"
	"mochilao-rentals/internal/rabbit"
	"mochilao-rentals/internal/rentals"

	"github.com/jasonlvhit/gocron"
)

func main() {
	connectionMongoDB, err := RegisterMongo()
	if err != nil {
		return
	}

	rentalsClient := database.NewRentalsClient(connectionMongoDB)
	rabbitmq, err := rabbit.ConnectRabbit()
	if err != nil {
		log.Fatal(err)
		return
	}

	rentals := rentals.NewGenerate(rentalsClient, rabbitmq)

	gocron.Every(3).Second().Do(rentals.GetRentalsAndDelete)

	// Start all the pending jobs
	<-gocron.Start()

}

func RegisterMongo() (mongo *database.Mongo, err error) {
	mongo = &database.Mongo{ConnectionString: config.Get().MongoConnectionString, DatabaseName: config.Get().DBName}
	// Make connection
	err = mongo.Connect(context.Background())

	return
}
