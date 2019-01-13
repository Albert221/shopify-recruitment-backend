package model

import "github.com/jinzhu/gorm"

type ProductOrder struct {
	gorm.Model
	Id        string `gorm:"primary_key"`
	ProductId string
	Product   func() *Product
	Quantity  int
}
