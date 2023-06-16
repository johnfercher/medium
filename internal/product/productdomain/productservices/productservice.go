package productservices

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"medium/m/v2/internal/product/productdb"
	"medium/m/v2/internal/product/productdomain/productentities"
)

type productService struct {
}

func New() *productService {
	return &productService{}
}

func (p *productService) GetByID(_ context.Context, id string) (*productentities.Product, error) {
	product, ok := productdb.Memory[id]
	if !ok {
		return nil, errors.New("product_not_found")
	}

	return product, nil
}

func (p *productService) Search(_ context.Context, productType string) ([]*productentities.Product, error) {
	var matchedValues []*productentities.Product
	for _, value := range productdb.Memory {
		if value.Type == productType {
			matchedValues = append(matchedValues, value)
		}
	}

	return matchedValues, nil
}

func (p *productService) Create(_ context.Context, product *productentities.Product) (*productentities.Product, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	idString := id.String()
	product.ID = idString
	productdb.Memory[idString] = product

	return product, nil
}

func (p *productService) Update(_ context.Context, productToUpdate *productentities.Product) (*productentities.Product, error) {
	product, ok := productdb.Memory[productToUpdate.ID]
	if !ok {
		err := errors.New("product_not_found")
		return nil, err
	}

	product.Type = productToUpdate.Type
	product.Quantity = productToUpdate.Quantity
	product.Name = productToUpdate.Name

	productdb.Memory[productToUpdate.ID] = product

	return product, nil
}

func (p *productService) Delete(_ context.Context, id string) error {
	productdb.Memory[id] = nil
	return nil
}
