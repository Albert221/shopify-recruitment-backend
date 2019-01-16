package model

type ProductOrder struct {
	Id        string `gorm:"primary_key"`
	ProductId string
	Product   Product
	Quantity  int
}
