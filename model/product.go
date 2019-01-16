package model

type Product struct {
	Id             string `gorm:"primary_key"`
	Title          string
	Price          float64
	InventoryCount int
}
