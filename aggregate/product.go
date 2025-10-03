package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nhloi9/ddd-go/entity"
)

var (
	//ErrProductNotFound is returned when a product is not found
	ErrProductNotFound = errors.New("the product was not found")
	//ErrProductAlreadyExist is returned when trying to add a product that already exists
	ErrProductAlreadyExist = errors.New("the product already exists")
	//ErrMissingValues is returned when trying to create a product with missing values
	ErrMissingValues = errors.New("missing values to create product")
)

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name, description string, price float64) (Product, error) {

	if name == "" || description == "" || price <= 0 {
		return Product{}, ErrMissingValues
	}
	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil

}

func (p *Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p *Product) GetItem() *entity.Item {
	return p.item
}

func (p *Product) GetPrice() float64 {
	return p.price
}
