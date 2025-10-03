package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/nhloi9/ddd-go/aggregate"
	"github.com/nhloi9/ddd-go/domain/customer"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

// Add implements customer.CustomerRepository.
func (m *MemoryRepository) Add(c aggregate.Customer) error {

	if m.customers == nil {
		m.Lock()
		m.customers = make(map[uuid.UUID]aggregate.Customer)
		m.Unlock()
	}

	if _, ok := m.customers[c.GetID()]; ok {
		return fmt.Errorf("customer with ID %s already exists", c.GetID().String())
	}
	m.Lock()
	m.customers[c.GetID()] = c
	m.Unlock()
	return nil

}

// Get implements customer.CustomerRepository.
func (m *MemoryRepository) Get(uuid uuid.UUID) (aggregate.Customer, error) {

	c, ok := m.customers[uuid]
	if !ok {
		return aggregate.Customer{}, customer.ErrCustomerNotFound
	}
	return c, nil
}

// Update implements customer.CustomerRepository.
func (m *MemoryRepository) Update(c aggregate.Customer) error {

	if _, ok := m.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer with ID %s does not exist", c.GetID().String())
	}
	m.Lock()
	m.customers[c.GetID()] = c
	m.Unlock()
	return nil
}

func New(customers map[uuid.UUID]aggregate.Customer) customer.CustomerRepository {
	return &MemoryRepository{
		customers: customers,
	}

}
