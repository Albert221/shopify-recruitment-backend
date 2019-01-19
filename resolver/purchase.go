package resolver

import (
	"github.com/Albert221/shopify-recruitment-backend/domain"
	"github.com/Albert221/shopify-recruitment-backend/domain/service"
	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

type PurchaseResolver struct {
	purchase *domain.Purchase
}

type PurchaseInput struct {
	CustomerName      string
	AddressFirstLine  string
	AddressSecondLine string
	City              string
	PostalCode        string
	Country           string

	CreditCardHolder  string
	CreditCardNumber  string
	CreditCardExpires int32
	CreditCardCVV     int32
}

type purchaseProductArgs struct {
	ProductId     string
	Quantity      int32
	PurchaseInput *PurchaseInput
}

func (r *RootResolver) PurchaseProduct(args purchaseProductArgs) (*PurchaseResolver, error) {
	// TODO: Validation, check if products are available!

	product := r.productsRepo.Get(args.ProductId)
	if product == nil {
		return nil, errors.New("given product does not exist")
	}

	charge := product.Price

	// *Get money from credit card*
	err := r.paymentGate.Charge(&service.CreditCardDetails{
		Holder:  args.PurchaseInput.CreditCardHolder,
		Number:  args.PurchaseInput.CreditCardNumber,
		Expires: int(args.PurchaseInput.CreditCardExpires),
		CVV:     int(args.PurchaseInput.CreditCardCVV),
	}, charge)
	if err != nil {
		return nil, errors.Wrap(err, "problem with payment has occurred")
	}

	// Create purchase object
	productOrder := domain.NewProductOrder(args.ProductId, int(args.Quantity))
	purchase := domain.NewPurchase([]domain.ProductOrder{*productOrder}, charge, &domain.Address{
		Name:       args.PurchaseInput.CustomerName,
		FirstLine:  args.PurchaseInput.AddressFirstLine,
		SecondLine: args.PurchaseInput.AddressSecondLine,
		City:       args.PurchaseInput.City,
		PostalCode: args.PurchaseInput.PostalCode,
		Country:    args.PurchaseInput.Country,
	})

	// Add purchase to db and remove quantities of products
	if err := r.purchaseRepo.Purchase(purchase); err != nil {
		return nil, errors.Wrap(err, "problem with purchase has occured")
	}

	return &PurchaseResolver{purchase: purchase}, nil
}

func (p *PurchaseResolver) Id() string {
	return p.purchase.Id
}

func (p *PurchaseResolver) Products() []*ProductOrderResolver {
	var resolvers []*ProductOrderResolver

	for _, productOrder := range p.purchase.Products {
		resolvers = append(resolvers, &ProductOrderResolver{productOrder: &productOrder})
	}

	return resolvers
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
	productOrder *domain.ProductOrder
}

func (o *ProductOrderResolver) Product() *ProductResolver {
	return &ProductResolver{product: &o.productOrder.Product}
}

func (o *ProductOrderResolver) Quantity() int32 {
	return int32(o.productOrder.Quantity)
}

// Address

type AddressResolver struct {
	address *domain.Address
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
