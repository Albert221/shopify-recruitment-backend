package model

import (
	"github.com/google/uuid"
	"time"
)

type Purchase struct {
	Id              string `gorm:"primary_key"`
	Products        []*ProductOrder
	Paid            float64
	ShippingAddress *Address
	PurchasedAt     time.Time
}

func NewPurchase(products []*ProductOrder, paid float64, shippingAddress *Address) *Purchase {
	return &Purchase{
		Id:              uuid.New().String(),
		Products:        products,
		Paid:            paid,
		ShippingAddress: shippingAddress,
		PurchasedAt:     time.Now(),
	}
}
