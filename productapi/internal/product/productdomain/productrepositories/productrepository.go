package productrepositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"medium/m/v2/internal/mysql"
	"medium/m/v2/internal/product/productdb"
	"medium/m/v2/internal/product/productdomain/productentities"
	"strings"
)

type ProductRepository struct {
	db *gorm.DB
}

func New() *ProductRepository {
	return &ProductRepository{
		db: mysql.DB,
	}
}

func (p *ProductRepository) GetByID(ctx context.Context, id string) (*productentities.Product, error) {
	product := &productentities.Product{}

	tx := p.db.Where("id = ?", id).First(product)

	if tx.Error != nil {
		return nil, errors.New("cannot_execute_query_error")
	}

	err := p.db.Model(product)
	if err != nil {
		return nil, errors.New("cannot_execute_query_error")
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
		return nil, errors.New("cannot_execute_query_error")
	}

	return products, nil
}

func (p *ProductRepository) Create(_ context.Context, product *productentities.Product) error {
	err := p.db.Transaction(func(tx *gorm.DB) error {
		if tx.Create(product).Error != nil {
			return errors.New("cannot_execute_query_error")
		}

		return nil
	})

	return err
}

func (p *ProductRepository) Update(_ context.Context, productToUpdate *productentities.Product) (*productentities.Product, error) {
	tx := p.db.Model(&productentities.Product{}).Where("id = ?", productToUpdate.ID).Updates(map[string]interface{}{
		"name":     productToUpdate.Name,
		"type":     productToUpdate.Type,
		"quantity": productToUpdate.Quantity,
	})

	if tx.Error != nil {
		return nil, errors.New("cannot_execute_query_error")
	}

	return nil
}

func (p *ProductRepository) Delete(_ context.Context, id string) error {
	productdb.Memory[id] = nil
	return nil
}
