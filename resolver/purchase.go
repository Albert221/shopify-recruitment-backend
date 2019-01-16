package resolver

import (
	"github.com/Albert221/shopify-recruitment-backend/model"
	"github.com/graph-gophers/graphql-go"
)

type PurchaseResolver struct {
	purchase *model.Purchase
}

type PurchaseInput struct {
	CustomerName      string
	AddressFirstLine  string
	AddressSecondLine *string
	City              string
	PostalCode        string
	Country           string

	CreditCardHolder  string
	CreditCardNumber  string
	CreditCardExpires string
	CreditCardCVV     int32
}

type purchaseProductArgs struct {
	ProductId     string
	Quantity      *int32
	PurchaseInput *PurchaseInput
}

func (c *RootResolver) PurchaseProduct(args purchaseProductArgs) (*PurchaseResolver, error) {
	return nil, nil
}

func (p *PurchaseResolver) Id() string {
	return p.purchase.Id
}

func (p *PurchaseResolver) Products() []*ProductOrderResolver {
	// todo
	return nil
}

func (p *PurchaseResolver) PurchasedAt() graphql.Time {
	return graphql.Time{Time: p.purchase.PurchasedAt}
}

func (p *PurchaseResolver) Paid() float64 {
	return p.purchase.Paid
}

func (p *PurchaseResolver) ShippingAddress() *AddressResolver {
	return &AddressResolver{address: p.purchase.ShippingAddress}
}

// ProductOrder

type ProductOrderResolver struct {
	productOrder *model.ProductOrder
}

func (o *ProductOrderResolver) Product() *ProductResolver {
	return &ProductResolver{product: &o.productOrder.Product}
}

func (o *ProductOrderResolver) Quantity() int32 {
	return int32(o.productOrder.Quantity)
}

// Address

type AddressResolver struct {
	address *model.Address
}

func (a *AddressResolver) Name() string {
	return a.address.Name
}

func (a *AddressResolver) AddressFirstLine() string {
	return a.address.FirstLine
}

func (a *AddressResolver) AddressSecondLine() *string {
	return &a.address.SecondLine
}

func (a *AddressResolver) City() string {
	return a.address.City
}

func (a *AddressResolver) PostalCode() string {
	return a.address.PostalCode
}

func (a *AddressResolver) Country() string {
	return a.address.Country
}
