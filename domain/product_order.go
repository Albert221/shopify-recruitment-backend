package domain

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type ProductOrder struct {
	Id         string `gorm:"primary_key"`
	ProductId  string
	Product    Product
	PurchaseId string
	Purchase   Purchase
	Quantity   int
}

func NewProductOrder(productId string, quantity int) *ProductOrder {
	return &ProductOrder{
		Id:        uuid.New().String(),
		ProductId: productId,
		Quantity:  quantity,
	}
}

func (ProductOrder) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("id", uuid.New().String())
	return nil
}