package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Id             string `gorm:"primary_key"`
	Title          string
	Price          Money
	InventoryCount int
}
