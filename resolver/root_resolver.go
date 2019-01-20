package resolver

import (
	"context"
	"github.com/Albert221/shopify-recruitment-backend/config"
	"github.com/Albert221/shopify-recruitment-backend/domain"
	"github.com/Albert221/shopify-recruitment-backend/domain/service"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/go-playground/validator.v9"
)

type RootResolver struct {
	productsRepo domain.ProductRepository
	purchaseRepo domain.PurchaseRepository
	cartRepo     domain.CartRepository
	paymentGate  service.PaymentGate
	cfg          *config.Configuration
	validator    *validator.Validate
}

type RootResolverArgs struct {
	ProductsRepo  domain.ProductRepository
	PurchaseRepo  domain.PurchaseRepository
	CartRepo      domain.CartRepository
	PaymentGate   service.PaymentGate
	Configuration *config.Configuration
}

func NewRootResolver(args *RootResolverArgs) *RootResolver {
	return &RootResolver{
		productsRepo: args.ProductsRepo,
		purchaseRepo: args.PurchaseRepo,
		cartRepo:     args.CartRepo,
		paymentGate:  args.PaymentGate,
		cfg:          args.Configuration,
		validator:    validator.New(),
	}
}

func (r *RootResolver) getClaims(ctx context.Context) jwt.MapClaims {
	token := ctx.Value("jwt")
	if token == nil {
		return nil
	}

	return ctx.Value("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
}
