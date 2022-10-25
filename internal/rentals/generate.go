package rentals

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"mochilao-rentals/internal/types"
)

type Rentals interface {
	Find3AndDelete() (result []types.Rental, err error)
}

type RabbitMq interface {
	Publish(body []byte) error
}

type Generate struct {
	Rentals  Rentals
	RabbitMq RabbitMq
}

func NewGenerate(rentals Rentals, rabbitMq RabbitMq) *Generate {
	return &Generate{
		Rentals:  rentals,
		RabbitMq: rabbitMq,
	}
}

func (gen *Generate) GetRentalsAndDelete() error {
	log.Print("Generate and send Rentals")
	rentals, err := gen.Rentals.Find3AndDelete()
	if err != nil {
		log.Print(err)
		return err
	}

	for idx, rental := range rentals {
		if idx == 2 {
			cities := []string{"Roma", "Paris", "Londres"}
			indx := rand.Intn(3)
			rental.Fields.City = cities[indx]
		}
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(rental)

		err = gen.RabbitMq.Publish(reqBodyBytes.Bytes())
		if err != nil {
			log.Print(err)
			return err
		}
	}
	return nil
}
