package model

import "time"

type Purchase struct {
	Id              string `gorm:"primary_key"`
	Products        []*ProductOrder
	Paid            float64
	ShippingAddress *Address
	PurchasedAt     time.Time
}
