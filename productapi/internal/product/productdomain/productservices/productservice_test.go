package productservices_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"medium/m/v2/internal/product/productdomain/productentities"
	"medium/m/v2/internal/product/productdomain/productrepositories/mocks"
	"medium/m/v2/internal/product/productdomain/productservices"
	"testing"
)

func TestNew(t *testing.T) {
	// Act
	sut := productservices.New(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*productservices.productService", fmt.Sprintf("%T", sut))
}

func TestProductService_Create_WhenCannotCreate_ShouldReturnAnError(t *testing.T) {
	// Arrange
	ctx := context.TODO()
	productToCreate := &productentities.Product{
		Name:     "name",
		Type:     "type",
		Quantity: 100,
	}

	errToReturn := errors.New("any_error")

	repository := mocks.NewProductRepository(t)
	repository.On("Create", ctx, productToCreate).Return(errToReturn)

	sut := productservices.New(repository)

	// Act
	product, err := sut.Create(ctx, productToCreate)

	// Assert
	assert.Nil(t, product)
	assert.Equal(t, errToReturn, err)
	repository.AssertNumberOfCalls(t, "Create", 1)
	repository.AssertCalled(t, "Create", ctx, productToCreate)
}

func TestProductService_Create_WhenEverythingWorks_ShouldReturnProduct(t *testing.T) {
	// Arrange
	ctx := context.TODO()
	productToCreate := &productentities.Product{
		Name:     "name",
		Type:     "type",
		Quantity: 100,
	}

	repository := mocks.NewProductRepository(t)
	repository.On("Create", ctx, productToCreate).Return(nil)

	sut := productservices.New(repository)

	// Act
	product, err := sut.Create(ctx, productToCreate)

	// Assert
	assert.Equal(t, productToCreate, product)
	assert.NotEmpty(t, product.ID)
	assert.Nil(t, err)
	repository.AssertNumberOfCalls(t, "Create", 1)
	repository.AssertCalled(t, "Create", ctx, productToCreate)
}
