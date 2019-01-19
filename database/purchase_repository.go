package database

import (
	"github.com/Albert221/shopify-recruitment-backend/domain"
	"github.com/jinzhu/gorm"
)

type PurchaseRepository struct {
	db *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) *PurchaseRepository {
	return &PurchaseRepository{db: db}
}

func (p *PurchaseRepository) Purchase(purchase *domain.Purchase) error {
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Save purchase
	if err := tx.Create(&purchase).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, order := range purchase.Products {
		// Update products inventory count
		var product domain.Product
		tx.Where("id = ?", order.ProductId).First(&product)
		newCount := product.InventoryCount - order.Quantity

		if err := tx.Model(&product).Update("inventory_count", newCount).Error; err != nil {
			tx.Rollback()
			return err
		}

		order.Product = product
	}

	// Query for updates purchase
	if err := tx.Preload("Products").
		Preload("Products.Product").First(&purchase).Error; err != nil {

		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
