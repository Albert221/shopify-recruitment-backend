package main

import (
	"context"
	"fmt"
	"github.com/Albert221/shopify-recruitment-backend/database"
	"github.com/Albert221/shopify-recruitment-backend/domain"
	"github.com/Albert221/shopify-recruitment-backend/resolver"
	"github.com/Albert221/shopify-recruitment-backend/stripe"
	"github.com/dgrijalva/jwt-go"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	db := createGorm()
	defer db.Close()

	// Resolver dependencies
	productsRepo := database.NewProductRepository(db)
	purchaseRepo := database.NewPurchaseRepository(db)
	cartRepo := database.NewCartRepository(db)
	stripeGate := stripe.PaymentGate{}

	rootResolver := resolver.NewRootResolver(productsRepo, purchaseRepo, cartRepo, stripeGate)

	schema := graphql.MustParseSchema(readSchema(), rootResolver)
	http.Handle("/api", jwtMiddleware(&relay.Handler{Schema: schema}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readSchema() string {
	data, err := ioutil.ReadFile("schema.graphql")

	if err != nil {
		panic(err)
	}

	return string(data)
}

func createGorm() *gorm.DB {
	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&domain.Product{}, &domain.ProductOrder{}, &domain.Purchase{},
		&domain.Cart{}, &domain.CartProductOrder{})

	return db.Debug()
}

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		authorization := r.Header.Get("Authorization")
		headerParts := strings.Split(authorization, " ")
		if len(headerParts) != 2 {
			next.ServeHTTP(w, r)
			return
		}
		jwtString := headerParts[1] // just after "Bearer"

		token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte("secret"), nil // FIXME: Use secret #1
		})
		if err != nil {
			fmt.Println(err)
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "jwt", token)))
	})
}
