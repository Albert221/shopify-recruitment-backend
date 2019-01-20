package resolver

import (
	"context"
	"errors"
	"github.com/Albert221/shopify-recruitment-backend/domain"
	"github.com/dgrijalva/jwt-go"
)

func (r *RootResolver) ShowCart(ctx context.Context) (*CartResolver, error) {
	cart, err := r.getCart(ctx)

	return &CartResolver{cart: cart}, err
}

func (r *RootResolver) CreateCart() (string, error) {
	cart := domain.NewCart()
	r.cartRepo.Save(cart)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cartId": cart.Id,
	})

	return token.SignedString(r.cfg.TokenSecret)
}

type addToCartArgs struct {
	ProductId string
	Quantity  int32 // fixme: validate if not negative
}

func (r *RootResolver) AddToCart(ctx context.Context, args addToCartArgs) (*CartResolver, error) {
	cart, err := r.getCart(ctx)
	if err != nil {
		return nil, err
	}

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

func (r *RootResolver) getCart(ctx context.Context) (*domain.Cart, error) {
	claims := r.getClaims(ctx)
	if claims == nil {
		return nil, errors.New("no cart token")
	}

	cartId, ok := claims["cartId"]
	if !ok {
		return nil, errors.New("invalid token, no cart")
	}

	cart := r.cartRepo.Get(cartId.(string))
	if cart == nil {
		return nil, errors.New("invalid token, no cart")
	}

	return cart, nil
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
