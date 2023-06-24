package productservices

import (
	"context"
	"github.com/google/uuid"
	"medium/m/v2/internal/product/productdomain/productentities"
	"medium/m/v2/internal/product/productdomain/productrepositories"
)

type productService struct {
	productRepository *productrepositories.ProductRepository
}

func New() *productService {
	return &productService{
		productRepository: productrepositories.New(),
	}
}

func (p *productService) GetByID(ctx context.Context, id string) (*productentities.Product, error) {
	return p.productRepository.GetByID(ctx, id)
}

func (p *productService) Search(ctx context.Context, productType string) ([]*productentities.Product, error) {
	return p.productRepository.Search(ctx, productType)
}

func (p *productService) Create(ctx context.Context, product *productentities.Product) (*productentities.Product, error) {
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

func (p *productService) Update(ctx context.Context, product *productentities.Product) (*productentities.Product, error) {
	err := p.productRepository.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productService) Delete(ctx context.Context, id string) error {
	return p.productRepository.Delete(ctx, id)
}
