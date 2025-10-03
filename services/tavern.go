package services

import (
	"log"

	"github.com/google/uuid"
)

type TavernConfiguration func(ts *Tavern) error

type Tavern struct {
	os *OrderService
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	// Create the orderservice
	ts := &Tavern{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(ts)
		if err != nil {
			return nil, err
		}
	}
	return ts, nil
}

func WithOrderService(os *OrderService) TavernConfiguration {
	return func(ts *Tavern) error {
		ts.os = os
		return nil
	}
}

// Order performs an order for a customer
func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.os.CreateOrder(customer, products)
	if err != nil {
		return err
	}
	log.Printf("Bill the Customer: %0.0f", price)

	// Bill the customer
	//err = t.BillingService.Bill(customer, price)
	return nil
}
