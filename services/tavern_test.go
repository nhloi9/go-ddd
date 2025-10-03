package services_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nhloi9/ddd-go/aggregate"
	"github.com/nhloi9/ddd-go/services"
)

func Test_MongoTavern(t *testing.T) {
	products := initProducts(t)
	customer, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	os, err := services.NewOrderService(
		// services.WithMongoCustomerRepository("mongodb://localhost:27017"),
		services.WithMemoryCustomerRepository([]aggregate.Customer{customer}),
		services.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := services.NewTavern(services.WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	// cust, err := aggregate.NewCustomer("Percy")
	// if err != nil {
	// 	t.Error(err)
	// }

	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(customer.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}
