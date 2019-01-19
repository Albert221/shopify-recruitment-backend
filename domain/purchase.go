package domain

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Purchase struct {
	Id              string `gorm:"primary_key"`
	Products        []ProductOrder
	Paid            float64
	ShippingAddress *Address `gorm:"embedded;embedded_prefix:address_"`
	PurchasedAt     time.Time
}

func NewPurchase(products []ProductOrder, paid float64, shippingAddress *Address) *Purchase {
	return &Purchase{
		Products:        products,
		Paid:            paid,
		ShippingAddress: shippingAddress,
		PurchasedAt:     time.Now(),
	}
}

func (Purchase) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("id", uuid.New().String())
	return nil
}