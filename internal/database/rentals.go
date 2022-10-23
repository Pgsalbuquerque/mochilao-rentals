package database

import (
	"context"
	"mochilao-rentals/internal/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RentalsClient struct {
	rentalsCollection *mongo.Collection
}

func NewRentalsClient(mongoConn *Mongo) *RentalsClient {
	return &RentalsClient{
		rentalsCollection: mongoConn.Collection(RentalsCollection),
	}
}

func (client *RentalsClient) Find3AndDelete() (result []types.Rental, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()
	for x := 0; x < 3; x++ {
		rental := types.Rental{}
		res := client.rentalsCollection.FindOneAndDelete(ctx, bson.M{})
		err = res.Decode(&rental)
		if err != nil {
			return nil, err
		}

		result = append(result, rental)
	}

	return
}
