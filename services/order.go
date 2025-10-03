// Package services holds all the services that connects repositories into a business flow
package services

import (
	"github.com/google/uuid"
	"github.com/nhloi9/ddd-go/aggregate"
	"github.com/nhloi9/ddd-go/domain/customer"
	mcr "github.com/nhloi9/ddd-go/domain/customer/memory"
	"github.com/nhloi9/ddd-go/domain/product"
	mpr "github.com/nhloi9/ddd-go/domain/product/memory"
)

// OrderConfiguration is an alias for a function that will take in a pointer to an OrderService and modify it
type OrderConfiguration func(os *OrderService) error

// OrderService is a implementation of the OrderService
type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

// NewOrderService takes a variable amount of OrderConfiguration functions and returns a new OrderService
// Each OrderConfiguration will be called in the order they are passed in
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	// Create the orderservice
	os := &OrderService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository applies a given customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithProductRepository(pr product.ProductRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.products = pr
		return nil
	}
}

// func WithMongoCustomerRepository(connectionString string) OrderConfiguration {
// 	return func(os *OrderService) error {
// 		// Create the mongo repo, if we needed parameters, such as connection strings they could be inputted here
// 		cr, err := mongo.New(context.Background(), connectionString)
// 		if err != nil {
// 			return err
// 		}
// 		os.customers = cr
// 		return nil
// 	}
// }

// WithMemoryCustomerRepository applies a memory customer repository to the OrderService
func WithMemoryCustomerRepository(customers []aggregate.Customer) OrderConfiguration {

	cr := mcr.New(map[uuid.UUID]aggregate.Customer{})
	for _, c := range customers {
		cr.Add(c)
	}

	return WithCustomerRepository(cr)
}

// WithMemoryProductRepository adds a in memory product repo and adds all input products
func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	pr := mpr.New()
	for _, p := range products {
		pr.Add(p)
	}
	return WithProductRepository(pr)
}

// CreateOrder will chaintogether all repositories to create a order for a customer
// will return the collected price of all Products
func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {

	_, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	total := 0.0

	for _, pid := range productIDs {

		p, err := o.products.GetByID(pid)
		if err != nil {
			return 0, err
		}
		total += p.GetPrice()
	}
	return total, nil
}
