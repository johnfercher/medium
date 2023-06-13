package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"medium/m/v2/internal/db"
	"medium/m/v2/internal/domain/entities"
)

type productService struct {
}

func NewProductService() *productService {
	return &productService{}
}

func (p *productService) GetByID(ctx context.Context, id string) (*entities.Product, error) {
	product, ok := db.Memory[id]
	if !ok {
		return nil, errors.New("product_not_found")
	}

	return product, nil
}

func (p *productService) Search(ctx context.Context, productType string) ([]*entities.Product, error) {
	var matchedValues []*entities.Product
	for _, value := range db.Memory {
		if value.Type == productType {
			matchedValues = append(matchedValues, value)
		}
	}

	return matchedValues, nil
}

func (p *productService) Create(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	idString := id.String()
	product.ID = idString
	db.Memory[idString] = product

	return product, nil
}

func (p *productService) Update(ctx context.Context, productToUpdate *entities.Product) (*entities.Product, error) {
	product, ok := db.Memory[productToUpdate.ID]
	if !ok {
		err := errors.New("product_not_found")
		return nil, err
	}

	product.Type = productToUpdate.Type
	product.Quantity = productToUpdate.Quantity
	product.Name = productToUpdate.Name

	db.Memory[productToUpdate.ID] = product

	return product, nil
}

func (p *productService) Delete(_ context.Context, id string) error {
	db.Memory[id] = nil
	return nil
}
