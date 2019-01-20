package database

import (
	"github.com/Albert221/shopify-recruitment-backend/domain"
	"github.com/jinzhu/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

// Get returns a Cart by specified id.
func (c *CartRepository) Get(id string) *domain.Cart {
	var cart domain.Cart
	c.db.Preload("Products").Preload("Products.Product").Where("id = ?", id).First(&cart)

	if cart.Id == "" {
		return nil
	}

	return &cart
}

// Save persists Cart to the database.
func (c *CartRepository) Save(cart *domain.Cart) error {
	return c.db.Save(cart).Error
}

