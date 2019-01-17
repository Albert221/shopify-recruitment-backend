package resolver

import "github.com/Albert221/shopify-recruitment-backend/model"

func (c *RootResolver) AllProducts(args struct{ OnlyAvailable *bool }) []*ProductResolver {
	var products []model.Product
	if args.OnlyAvailable != nil && *args.OnlyAvailable {
		c.db.Where("inventory_count > 0").Find(&products)
	} else {
		c.db.Find(&products)
	}

	var productResolvers []*ProductResolver
	for _, product := range products {
		productResolvers = append(productResolvers, &ProductResolver{product: &product})
	}

	return productResolvers
}

func (c *RootResolver) Product(args struct{ ProductId string }) *ProductResolver {
	var product model.Product
	c.db.First(&product, "id = ?", args.ProductId)

	if product.Id == "" {
		return nil
	}

	return &ProductResolver{product: &product}
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

func (p *ProductResolver) Price() float64 {
	return p.product.Price
}

func (p *ProductResolver) InventoryCount() int32 {
	return int32(p.product.InventoryCount)
}
