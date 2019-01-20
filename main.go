package main

import (
	"context"
	"fmt"
	"github.com/Albert221/shopify-recruitment-backend/config"
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
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("there was an error with configuration file: %s", err)
	}

	apiHandler, dbClose, err := createApiHandler(cfg)
	if err != nil {
		log.Fatalf("there was an error with api handler: %s", err)
	}
	defer dbClose()

	http.Handle(cfg.ApiPath, apiHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.HttpPort), nil))
}

func createApiHandler(cfg *config.Configuration) (http.Handler, func(), error) {
	db, err := createGorm(cfg.Debug, cfg.DbFile)
	if err != nil {
		return nil, nil, err
	}

	// Resolver dependencies
	rootResolverArgs := &resolver.RootResolverArgs{
		ProductsRepo:  database.NewProductRepository(db),
		PurchaseRepo:  database.NewPurchaseRepository(db),
		CartRepo:      database.NewCartRepository(db),
		PaymentGate:   stripe.NewPaymentGate(),
		Configuration: cfg,
	}

	rootResolver := resolver.NewRootResolver(rootResolverArgs)

	sch, err := readSchema()
	if err != nil {
		return nil, nil, err
	}

	schema := graphql.MustParseSchema(sch, rootResolver)
	handler := jwtMiddleware(&relay.Handler{Schema: schema}, cfg.TokenSecret)

	return handler, func() { db.Close() }, nil
}

func readSchema() (string, error) {
	data, err := ioutil.ReadFile("schema.graphql")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func createGorm(debug bool, dbFile string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&domain.Product{}, &domain.ProductOrder{}, &domain.Purchase{},
		&domain.Cart{}, &domain.CartProductOrder{})

	if debug {
		return db.Debug(), nil
	}

	return db, nil
}

func jwtMiddleware(next http.Handler, secret []byte) http.Handler {
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

			return secret, nil
		})
		if err != nil {
			fmt.Println(err)
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "jwt", token)))
	})
}
