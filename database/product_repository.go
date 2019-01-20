package database

import (
	"github.com/Albert221/shopify-recruitment-backend/domain"
	"github.com/jinzhu/gorm"
	"strings"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Get returns a Product by specified id or nil if it does not exist.
func (p *ProductRepository) Get(id string) *domain.Product {
	var product domain.Product
	p.db.First(&product, "id = ?", id)

	if product.Id == "" {
		return nil
	}

	return &product
}

// GetMany returns a slice of Products by specified ids.
func (p *ProductRepository) GetMany(ids []string) []*domain.Product {
	var products []*domain.Product

	stmts := strings.TrimRight(strings.Repeat("?, ", len(ids)), ", ")
	params := make([]interface{}, len(stmts))
	for i, id := range ids {
		params[i] = id
	}

	p.db.Where("id IN (" + stmts + ")", params...).Find(&products)

	return products
}

// GetAll returns a slice of all Products from database
func (p *ProductRepository) GetAll() []*domain.Product {
	var products []*domain.Product
	p.db.Find(&products)

	return products
}

// GetAvailable returns a slice of all Products that have positive inventory count.
func (p *ProductRepository) GetAvailable() []*domain.Product {
	var products []*domain.Product
	p.db.Where("inventory_count > 0").Find(&products)

	return products
}

// Save persists Product to the database.
func (p *ProductRepository) Save(product *domain.Product) error {
	return p.db.Save(product).Error
}
