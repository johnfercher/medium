package productrepositories

import (
	"context"
	"gorm.io/gorm"
	"medium/m/v2/internal/product/productdomain/productentities"
	"strings"
)

type ProductRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) GetByID(_ context.Context, id string) (*productentities.Product, error) {
	product := &productentities.Product{}

	tx := p.db.Where("id = ?", id).First(product)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return product, nil
}

func (p *ProductRepository) Search(_ context.Context, productType string) ([]*productentities.Product, error) {
	limit := 100
	products := []*productentities.Product{}

	query := []string{}
	args := []interface{}{}

	query = append(query, "products.type = ?")
	args = append(args, productType)

	tx := p.db.Table("products").
		Select("products.id, products.name, products.type, products.quantity")

	tx = tx.Where(strings.Join(query, " AND "), args...)

	tx = tx.Limit(limit)

	tx = tx.Scan(&products)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return products, nil
}

func (p *ProductRepository) Create(_ context.Context, product *productentities.Product) error {
	tx := p.db.Create(product)
	return tx.Error
}

func (p *ProductRepository) Update(_ context.Context, productToUpdate *productentities.Product) error {
	tx := p.db.Model(&productentities.Product{}).Where("id = ?", productToUpdate.ID).Updates(map[string]interface{}{
		"name":     productToUpdate.Name,
		"type":     productToUpdate.Type,
		"quantity": productToUpdate.Quantity,
	})

	return tx.Error
}

func (p *ProductRepository) Delete(_ context.Context, id string) error {
	tx := p.db.Where("id = ?", id).Delete(&productentities.Product{})
	return tx.Error
}
