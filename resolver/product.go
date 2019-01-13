package resolver

import "github.com/Albert221/shopify-recruitment-backend/model"

func (c *RootResolver) AllProducts(args struct { OnlyAvailable *bool }) []*ProductResolver {
	return nil
}

func (c *RootResolver) Product(args struct { ProductId string }) (*ProductResolver, error) {
	return nil, nil
}

type ProductResolver struct {
	product *model.Product
}

func (p *ProductResolver) Id() string {
	return p.product.Id
}

func (p *ProductResolver) Title() string {
	return p.product.Title
}

func (p *ProductResolver) Price() Money {
	return Money{p.product.Price}
}

func (p *ProductResolver) InventoryCount() int32 {
	return int32(p.product.InventoryCount)
}