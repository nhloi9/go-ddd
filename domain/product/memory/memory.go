package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/nhloi9/ddd-go/aggregate"
	"github.com/nhloi9/ddd-go/domain/product"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

// Add implements product.ProductRepository.
func (m *MemoryProductRepository) Add(p aggregate.Product) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[p.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	m.products[p.GetID()] = p
	return nil
}

// Delete implements product.ProductRepository.
func (m *MemoryProductRepository) Delete(id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(m.products, id)
	return nil
}

// GetAll implements product.ProductRepository.
func (m *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {

	products := make([]aggregate.Product, len(m.products))

	for _, p := range m.products {
		products = append(products, p)
	}
	return products, nil

}

// GetByID implements product.ProductRepository.
func (m *MemoryProductRepository) GetByID(id uuid.UUID) (p aggregate.Product, e error) {

	p, ok := m.products[id]
	if !ok {
		e = product.ErrProductNotFound
	}
	return
}

// Update implements product.ProductRepository.
func (m *MemoryProductRepository) Update(p aggregate.Product) error {

	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[p.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	m.products[p.GetID()] = p

	return nil

}

func New() product.ProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}
