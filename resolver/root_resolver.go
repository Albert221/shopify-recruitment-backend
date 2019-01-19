package resolver

import (
	"github.com/Albert221/shopify-recruitment-backend/domain"
	"github.com/Albert221/shopify-recruitment-backend/domain/service"
)

type RootResolver struct {
	productsRepo domain.ProductRepository
	purchaseRepo domain.PurchaseRepository
	paymentGate  service.PaymentGate
}

func NewRootResolver(productsRepo domain.ProductRepository, purchaseRepo domain.PurchaseRepository,
	paymentGate service.PaymentGate) *RootResolver {
	return &RootResolver{productsRepo: productsRepo, purchaseRepo: purchaseRepo, paymentGate: paymentGate}
}
