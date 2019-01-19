package domain

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type CartProductOrder struct {
	Id        string `gorm:"primary_key"`
	ProductId string
	Product   Product
	CartId    string
	Cart      Cart
	Quantity  int
}

func NewCartProductOrder(productId string, quantity int) *CartProductOrder {
	return &CartProductOrder{
		ProductId: productId,
		Quantity:  quantity,
	}
}

func (CartProductOrder) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("id", uuid.New().String())
	return nil
}
