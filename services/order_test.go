package services_test

import (
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/nhloi9/ddd-go/aggregate"
	"github.com/nhloi9/ddd-go/services"
)

func initProducts(t *testing.T) []aggregate.Product {
	var products []aggregate.Product
	for i := 0; i < 100; i++ {
		p, err := aggregate.NewProduct("product"+strconv.Itoa(i), "description", float64(i)*1.1+1)
		if err != nil {
			t.Error(err)
		}
		products = append(products, p)
	}
	return products
}

func TestOrderService_CreateOrder(t *testing.T) {

	customer, err := aggregate.NewCustomer("pen")
	if err != nil {
		t.Fatalf("failed to create customer: %v", err)
	}

	products := initProducts(t)

	orderService, err := services.NewOrderService(services.WithMemoryCustomerRepository([]aggregate.Customer{customer}), services.WithMemoryProductRepository(products))
	if err != nil {
		t.Fatalf("failed to create order service: %v", err)
	}

	// Assert
	totalPrice, err := orderService.CreateOrder(customer.GetID(), []uuid.UUID{products[0].GetID(), products[1].GetID(), products[2].GetID()})
	if err != nil {
		t.Fatalf("failed to create order: %v", err)
	}

	expectedTotal := products[0].GetPrice() + products[1].GetPrice() + products[2].GetPrice()
	if totalPrice != expectedTotal {
		t.Errorf("expected total price %v, got %v", expectedTotal, totalPrice)
	}
}
