package domain

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Product struct {
	Id             string `gorm:"primary_key"`
	Title          string
	Price          float64
	InventoryCount int
}

func (Product) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("id", uuid.New().String())
	return nil
}