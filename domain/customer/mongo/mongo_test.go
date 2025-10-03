package mongo_test

import (
	"context"
	"testing"

	"github.com/nhloi9/ddd-go/aggregate"
	"github.com/nhloi9/ddd-go/domain/customer/mongo"
	mg "go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func TestMongo_GetCustomer(t *testing.T) {
	t.Error("This test is expected to fail until a real MongoDB instance is available.")
	t.Log("This message will only appear if the test fails or -v is used.")
	t.Logf("Value: %d", 42)

	client, _ := mg.Connect(options.Client().ApplyURI("mongodb+srv://nguyenhuuloi:MX1TSfgbbady2e6f@cluster0.byr8s5v.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())
	collection := client.Database("go-ddd").Collection("customers")
	repo := mongo.NewMongoCustomerRepository(collection)

	customer, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Printf("Customer ID: %s\n", customer.GetID().String())
	t.Logf("Customer ID: %s\n", customer.GetID().String())

	err = repo.Add(customer)
	if err != nil {
		t.Fatal(err)
	}

	got, err := repo.Get(customer.GetID())
	if err != nil {
		t.Fatal(err)
	}

	if got.GetID() != customer.GetID() {
		t.Errorf("Expected ID %s, got %s", customer.GetID(), got.GetID())
	}
}
