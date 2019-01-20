package domain

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Cart struct {
	Id       string `gorm:"primary_key"`
	Products []CartProductOrder
}

func NewCart() *Cart {
	return &Cart{}
}

// Total returns a total amount of money for products.
func (c *Cart) Total() float64 {
	total := 0.0
	for _, order := range c.Products {
		total += order.Product.Price * float64(order.Quantity)
	}

	return total
}

func (Cart) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("id", uuid.New().String())
	return nil
}