package mongo

import (
	"context"

	"github.com/google/uuid"
	"github.com/nhloi9/ddd-go/aggregate"
	"github.com/nhloi9/ddd-go/domain/customer"

	"go.mongodb.org/mongo-driver/v2/bson"
	mg "go.mongodb.org/mongo-driver/v2/mongo"
)

// var client, _ = mg.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))

type MongoCustomerRepository struct {
	collection *mg.Collection
}

type MongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

// NewFromCustomer takes in a aggregate and converts into internal structure
func NewFromCustomer(c aggregate.Customer) MongoCustomer {
	return MongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

// ToAggregate converts into a aggregate.Customer
// this could validate all values present etc
func (m MongoCustomer) ToAggregate() aggregate.Customer {
	c := aggregate.Customer{}
	c.SetID(m.ID)
	c.SetName(m.Name)

	return c

}

// Add implements customer.CustomerRepository.
func (m *MongoCustomerRepository) Add(c aggregate.Customer) error {
	_, err := m.collection.InsertOne(context.Background(), NewFromCustomer(c))
	return err
}

// Get implements customer.CustomerRepository.
func (m *MongoCustomerRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	var c MongoCustomer
	err := m.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&c)
	return c.ToAggregate(), err
}

// Update implements customer.CustomerRepository.
func (m *MongoCustomerRepository) Update(aggregate.Customer) error {
	panic("unimplemented")
}

func NewMongoCustomerRepository(
	collection *mg.Collection,
) customer.CustomerRepository {
	return &MongoCustomerRepository{
		collection: collection,
	}
}
