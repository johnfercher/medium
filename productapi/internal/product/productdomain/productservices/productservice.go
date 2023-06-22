package productservices

import (
	"context"
	"github.com/google/uuid"
	"medium/m/v2/internal/product/productdomain/productentities"
	"medium/m/v2/internal/product/productdomain/productrepositories"
)

type ProductService struct {
	productRepository *productrepositories.ProductRepository
}

func New(productRepository *productrepositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (p *ProductService) GetByID(ctx context.Context, id string) (*productentities.Product, error) {
	return p.productRepository.GetByID(ctx, id)
}

func (p *ProductService) Search(ctx context.Context, productType string) ([]*productentities.Product, error) {
	return p.productRepository.Search(ctx, productType)
}

func (p *ProductService) Create(ctx context.Context, product *productentities.Product) (*productentities.Product, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	idString := id.String()
	product.ID = idString

	err = p.productRepository.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) Update(ctx context.Context, product *productentities.Product) (*productentities.Product, error) {
	err := p.productRepository.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) Delete(ctx context.Context, id string) error {
	return p.productRepository.Delete(ctx, id)
}
