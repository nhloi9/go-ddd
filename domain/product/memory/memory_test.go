package memory_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nhloi9/ddd-go/aggregate"
	"github.com/nhloi9/ddd-go/domain/product"
	"github.com/nhloi9/ddd-go/domain/product/memory"
)

func TestMemoryProductRepository_Add(t *testing.T) {
	repo := memory.New()
	product, err := aggregate.NewProduct("Beer", "Good for you're health", 1.99)
	if err != nil {
		t.Error(err)
	}

	err = repo.Add(product)
	if err != nil {
		t.Error(err)
	}
}
func TestMemoryProductRepository_Get(t *testing.T) {
	repo := memory.New()
	existingProd, err := aggregate.NewProduct("Beer", "Good for you're health", 1.99)
	if err != nil {
		t.Error(err)
	}

	repo.Add(existingProd)

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get product by id",
			id:          existingProd.GetID(),
			expectedErr: nil,
		}, {
			name:        "Get non-existing product by id",
			id:          uuid.New(),
			expectedErr: product.ErrProductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetByID(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}

}
func TestMemoryProductRepository_Delete(t *testing.T) {
	repo := memory.New()
	existingProd, err := aggregate.NewProduct("Beer", "Good for you're health", 1.99)
	if err != nil {
		t.Error(err)
	}

	repo.Add(existingProd)

	err = repo.Delete(existingProd.GetID())
	if err != nil {
		t.Error(err)
	}

}
