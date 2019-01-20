package resolver

import (
	"context"
	"github.com/Albert221/shopify-recruitment-backend/domain"
	"github.com/Albert221/shopify-recruitment-backend/domain/service"
	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

type PurchaseResolver struct {
	purchase *domain.Purchase
}

type PurchaseInput struct {
	CustomerName      string `validate:"required"`
	AddressFirstLine  string `validate:"required"`
	AddressSecondLine string
	City              string `validate:"required"`
	PostalCode        string `validate:"required"`
	Country           string `validate:"required"`

	CreditCardHolder  string `validate:"required"`
	CreditCardNumber  string `validate:"required,len=19,numeric"`
	CreditCardExpires string `validate:"required,len=4"`
	CreditCardCVV     string `validate:"required,len=3"`
}

type purchaseProductArgs struct {
	ProductId     string
	Quantity      int32 `validate:"min=1"`
	PurchaseInput *PurchaseInput
}

func (r *RootResolver) PurchaseProduct(args purchaseProductArgs) (*PurchaseResolver, error) {
	err := r.validator.Struct(args)
	if err != nil {
		return nil, err
	}

	product := r.productsRepo.Get(args.ProductId)
	if product == nil {
		return nil, errors.New("given product does not exist")
	}

	charge := product.Price
	productIdQuantity := map[string]int{args.ProductId: int(args.Quantity)}

	return r.purchaseProducts(args.PurchaseInput, productIdQuantity, charge)
}

func (r *RootResolver) CheckoutCart(ctx context.Context, args struct{ PurchaseInput PurchaseInput }) (*PurchaseResolver, error) {
	err := r.validator.Struct(args)
	if err != nil {
		return nil, err
	}

	cart, err := r.getCart(ctx)
	if err != nil {
		return nil, err
	}

	charge := 0.0
	productIdQuantity := make(map[string]int)
	for _, order := range cart.Products {
		charge += order.Product.Price * float64(order.Quantity)
		productIdQuantity[order.ProductId] = order.Quantity
	}

	return r.purchaseProducts(&args.PurchaseInput, productIdQuantity, charge)
}

func (r *RootResolver) purchaseProducts(input *PurchaseInput, productIdQuantity map[string]int, charge float64) (*PurchaseResolver, error) {
	if !r.productsAvailable(productIdQuantity) {
		return nil, errors.New("some of the products are not available. check their inventoryCount")
	}

	// *Get money from credit card*
	err := r.paymentGate.Charge(&service.CreditCardDetails{
		Holder:  input.CreditCardHolder,
		Number:  input.CreditCardNumber,
		Expires: input.CreditCardExpires,
		CVV:     input.CreditCardCVV,
	}, charge)
	if err != nil {
		return nil, errors.Wrap(err, "problem with payment has occurred")
	}

	// Create purchase object
	var productOrders []domain.ProductOrder
	for productId, quantity := range productIdQuantity {
		productOrders = append(productOrders, *domain.NewProductOrder(productId, quantity))
	}

	purchase := domain.NewPurchase(productOrders, charge, &domain.Address{
		Name:       input.CustomerName,
		FirstLine:  input.AddressFirstLine,
		SecondLine: input.AddressSecondLine,
		City:       input.City,
		PostalCode: input.PostalCode,
		Country:    input.Country,
	})

	// Add purchase to db and remove quantities of products
	if err := r.purchaseRepo.Purchase(purchase); err != nil {
		return nil, errors.Wrap(err, "problem with purchase has occured")
	}

	return &PurchaseResolver{purchase: purchase}, nil
}

func (r *RootResolver) productsAvailable(productIdQuantity map[string]int) bool {
	var productIds []string
	for id := range productIdQuantity {
		productIds = append(productIds, id)
	}

	products := r.productsRepo.GetMany(productIds)
	if len(products) != len(productIdQuantity) {
		return false
	}
	for _, product := range products {
		quantity, exists := productIdQuantity[product.Id]
		if !exists || quantity > product.InventoryCount {
			return false
		}
	}

	return true
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
