package productrepositories

import (
	"context"
	"errors"
	"medium/m/v2/internal/product/productdb"
	"medium/m/v2/internal/product/productdomain/productentities"
)

type ProductRepository struct {
}

func New() *ProductRepository {
	return &ProductRepository{}
}

func (p *ProductRepository) GetByID(_ context.Context, id string) (*productentities.Product, error) {
	product, ok := productdb.Memory[id]
	if !ok {
		return nil, errors.New("product_not_found")
	}

	return product, nil
}

func (p *ProductRepository) Search(_ context.Context, productType string) ([]*productentities.Product, error) {
	var matchedValues []*productentities.Product
	for _, value := range productdb.Memory {
		if value.Type == productType {
			matchedValues = append(matchedValues, value)
		}
	}

	return matchedValues, nil
}

func (p *ProductRepository) Create(ctx context.Context, product *productentities.Product) error {
	productdb.Memory[product.ID] = product
	return nil
}

func (p *ProductRepository) Update(_ context.Context, productToUpdate *productentities.Product) error {
	product, ok := productdb.Memory[productToUpdate.ID]
	if !ok {
		return errors.New("product_not_found")
	}

	product.Type = productToUpdate.Type
	product.Quantity = productToUpdate.Quantity
	product.Name = productToUpdate.Name

	productdb.Memory[productToUpdate.ID] = product

	return nil
}

func (p *ProductRepository) Delete(_ context.Context, id string) error {
	productdb.Memory[id] = nil
	return nil
}
