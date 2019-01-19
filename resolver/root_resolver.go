package resolver

import (
	"context"
	"github.com/Albert221/shopify-recruitment-backend/domain"
	"github.com/Albert221/shopify-recruitment-backend/domain/service"
	"github.com/dgrijalva/jwt-go"
)

type RootResolver struct {
	productsRepo domain.ProductRepository
	purchaseRepo domain.PurchaseRepository
	cartRepo     domain.CartRepository
	paymentGate  service.PaymentGate
}

func NewRootResolver(productsRepo domain.ProductRepository, purchaseRepo domain.PurchaseRepository,
	cartRepo domain.CartRepository, paymentGate service.PaymentGate) *RootResolver {
	return &RootResolver{productsRepo: productsRepo, purchaseRepo: purchaseRepo, cartRepo: cartRepo,
		paymentGate: paymentGate}
}

func (r *RootResolver) getClaims(ctx context.Context) jwt.MapClaims {
	token := ctx.Value("jwt")
	if token == nil {
		return nil
	}

	return ctx.Value("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
}
