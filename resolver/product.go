package resolver

import (
	"github.com/Albert221/shopify-recruitment-backend/domain"
)

func (r *RootResolver) AllProducts(args struct{ OnlyAvailable *bool }) []*ProductResolver {
	var products []*domain.Product
	if args.OnlyAvailable != nil && *args.OnlyAvailable {
		products = r.productsRepo.GetAvailable()
	} else {
		products = r.productsRepo.GetAll()
	}

	var productResolvers []*ProductResolver
	for _, product := range products {
		productResolvers = append(productResolvers, &ProductResolver{product: product})
	}

	return productResolvers
}

func (r *RootResolver) Product(args struct{ ProductId string }) *ProductResolver {
	product := r.productsRepo.Get(args.ProductId)

	if product == nil {
		return nil
	}

	return &ProductResolver{product: product}
}

type ProductResolver struct {
	product *domain.Product
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
