package model

import (
	"github.com/jinzhu/gorm"
)

type Purchase struct {
	gorm.Model
	Id              string `gorm:"primary_key"`
	Products        []*ProductOrder
	Paid            Money
	ShippingAddress *Address
}
