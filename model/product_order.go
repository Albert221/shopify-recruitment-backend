package model

import "github.com/google/uuid"

type ProductOrder struct {
	Id        string `gorm:"primary_key"`
	ProductId string
	Product   Product
	Quantity  int
}

func NewProductOrder(productId string, quantity int) *ProductOrder {
	return &ProductOrder{
		Id:        uuid.New().String(),
		ProductId: productId,
		Quantity:  quantity,
	}
}
