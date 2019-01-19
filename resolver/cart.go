package resolver

import (
	"context"
	"errors"
	"github.com/Albert221/shopify-recruitment-backend/domain"
	"github.com/dgrijalva/jwt-go"
)

func (r *RootResolver) ShowCart(ctx context.Context) (*CartResolver, error) {
	claims := r.getClaims(ctx)
	if claims == nil {
		return nil, errors.New("user unauthorized")
	}

	cartId, ok := claims["cartId"]
	if !ok {
		return nil, nil
	}

	cart := r.cartRepo.Get(cartId.(string))
	if cart == nil {
		return nil, errors.New("invalid token, no cart")
	}

	return &CartResolver{cart: cart}, nil
}

func (r *RootResolver) CreateCart() (string, error) {
	cart := domain.NewCart()
	r.cartRepo.Save(cart)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cartId": cart.Id,
	})

	return token.SignedString([]byte("secret")) // FIXME: Use secret #2
}

type addToCartArgs struct {
	ProductId string
	Quantity  int32
}

func (r *RootResolver) AddToCart(ctx context.Context, args addToCartArgs) (*CartResolver, error) {
	claims := r.getClaims(ctx)
	if claims == nil {
		return nil, errors.New("user unauthorized")
	}

	cartId, ok := claims["cartId"]
	if !ok {
		return nil, nil
	}

	cart := r.cartRepo.Get(cartId.(string))
	if cart == nil {
		return nil, errors.New("invalid token, no cart")
	}
	// FIXME: Keep stuff above DRY!

	product := r.productsRepo.Get(args.ProductId)
	if product == nil {
		return nil, errors.New("given product does not exist")
	}

	alreadyInCart := false
	for i, order := range cart.Products {
		if order.ProductId == args.ProductId {
			alreadyInCart = true
			cart.Products[i].Quantity += int(args.Quantity)
			break
		}
	}

	if !alreadyInCart {
		order := domain.NewCartProductOrder(args.ProductId, int(args.Quantity))
		cart.Products = append(cart.Products, *order)
	}

	r.cartRepo.Save(cart)
	cart = r.cartRepo.Get(cart.Id)

	return &CartResolver{cart: cart}, nil
}

type CartResolver struct {
	cart *domain.Cart
}

func (c *CartResolver) Products() []*CartProductOrderResolver {
	var resolvers []*CartProductOrderResolver
	for _, order := range c.cart.Products {
		resolvers = append(resolvers, &CartProductOrderResolver{productOrder: &order})
	}

	return resolvers
}

func (c *CartResolver) Total() float64 {
	return c.cart.Total()
}

type CartProductOrderResolver struct {
	productOrder *domain.CartProductOrder
}

func (c *CartProductOrderResolver) Product() *ProductResolver {
	return &ProductResolver{product: &c.productOrder.Product}
}

func (c *CartProductOrderResolver) Quantity() int32 {
	return int32(c.productOrder.Quantity)
}
