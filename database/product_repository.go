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

func (p *ProductRepository) Get(id string) *domain.Product {
	var product domain.Product
	p.db.First(&product, "id = ?", id)

	return &product
}

func (p *ProductRepository) GetMany(ids []string) []*domain.Product {
	var products []*domain.Product

	stmts := strings.TrimRight(strings.Repeat("?, ", len(ids)), ", ")
	params := make([]interface{}, len(stmts))
	for i, stmt := range stmts {
		params[i] = stmt
	}

	p.db.Where("id IN (" + stmts + ")", params...).Find(&products)

	return products
}

func (p *ProductRepository) GetAll() []*domain.Product {
	var products []*domain.Product
	p.db.Find(&products)

	return products
}

func (p *ProductRepository) GetAvailable() []*domain.Product {
	var products []*domain.Product
	p.db.Where("inventory_count > 0").Find(&products)

	return products
}

func (p *ProductRepository) Save(product *domain.Product) error {
	return p.db.Save(product).Error
}
