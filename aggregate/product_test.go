package aggregate_test

import (
	"testing"

	"github.com/nhloi9/ddd-go/aggregate"
)

func TestProduct_NewProduct(t *testing.T) {

	type testCase struct {
		name               string
		productName        string
		productDescription string
		price              float64
		expectedErr        error
	}

	testCases := []testCase{
		{
			name:               "Valid Product",
			productName:        "Product 1",
			productDescription: "This is product 1",
			price:              10.0,
			expectedErr:        nil,
		}, {
			name:               "Missing Name",
			productName:        "",
			productDescription: "This is product 1",
			price:              10.0,
			expectedErr:        aggregate.ErrMissingValues,
		}, {
			name:               "Missing Description",
			productName:        "Product 1",
			productDescription: "",
			price:              10.0,
			expectedErr:        aggregate.ErrMissingValues,
		}, {
			name:               "Zero Price",
			productName:        "Product 1",
			productDescription: "This is product 1",
			price:              0.0,
			expectedErr:        aggregate.ErrMissingValues,
		}, {
			name:               "Negative Price",
			productName:        "Product 1",
			productDescription: "This is product 1",
			price:              -10.0,
			expectedErr:        aggregate.ErrMissingValues,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			_, err := aggregate.NewProduct(tc.productName, tc.productDescription, tc.price)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}

}
