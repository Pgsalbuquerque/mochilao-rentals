package rentals

import (
	"bytes"
	"encoding/json"
	"log"
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
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(rentals)

	err = gen.RabbitMq.Publish(reqBodyBytes.Bytes())
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
