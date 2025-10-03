package aggregate_test

import (
	"testing"

	"github.com/nhloi9/ddd-go/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{{
		test:        "test new customer with valid name",
		name:        "John Doe",
		expectedErr: nil,
	},
		{
			test:        "test new customer with empty name",
			name:        "",
			expectedErr: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
