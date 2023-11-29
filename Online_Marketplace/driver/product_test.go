package driver

import (
	"fmt"
	"onlinemarketplace/model"
	"testing"

	"errors"

	"github.com/stretchr/testify/assert"
)

// Sample testing a hypothetical `CreateProduct` function
func TestCreateProduct(t *testing.T) {
	NewDBConn(true)
	// Set up test cases with both valid and invalid input data
	cases := []struct {
		input        model.Product
		success      bool
		errorMessage error
	}{
		{
			model.Product{Name: "MI", Description: "Redmi", Price: 20000.0}, true, nil,
		},
		{
			model.Product{Name: "", Description: "Samsung", Price: 10.0}, false, errors.New("Name shouldn't be empty"),
		},
		{
			model.Product{Name: "Moto G", Description: "", Price: 10.0}, false, errors.New("Description shouldn't be empty"),
		},
		{
			model.Product{Name: "Moto G10", Description: "Moto"}, false, errors.New("Price shouldn't be zero or negative"),
		},
		{
			model.Product{Name: "Samsung M11", Description: "Samsung", Price: -10.0}, false, errors.New("Price shouldn't be zero or negative"),
		},
		{
			model.Product{}, false, errors.New("Name shouldn't be empty"),
		},
		{
			model.Product{Name: "Samsung M11", Description: "E3MDEyMzM4NzMsInBhc3N3b3JkIjoicGFzcyIsInVzZXJuYW1lIjoidXNlcjA1In0.Lx95rNQZjGh4RaKwf4tiTcQTquzf1zs13Os7rFdZ_xweyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDEyMzM4NzMsInBhc3N3b3JkIjoicGFzcyIsInVzZXJuYW1lIjoidXNlcjA1In0.Lx95rNQZjGh4RaKwf4tiTcQTquzf1zs13Os7rFdZ_xw", Price: 10.0}, false, errors.New("Description shouldn't be greater than 200 character"),
		},
		{
			model.Product{Name: "E3MDEyMzM4NzMsInBhc3N3b3JkIjoicGFzcyIsInRaKwf4tiTcQTquzf1zs13Os7rFI6IkpXVCJ9eyJleHAiOjE3MDEyMzM4NzMsInBNlcjA1In0", Description: "Samsung", Price: 10.0}, false, errors.New("Name shouldn't be greater than 100 character"),
		},
	}

	for i, tc := range cases {
		err := CreateProduct(&tc.input)

		if tc.success {
			if assert.Nil(t, err) {
				fmt.Printf("Test case %d Passed %v\n", i, err)
			}
		} else {
			if assert.Equal(t, err, tc.errorMessage) {
				fmt.Printf("Test case %d Passed %v\n", i, err)
			}
		}
	}
	GetDBConn().Exec("DELETE FROM products")
}
