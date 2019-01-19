package main

import (
	"github.com/Albert221/shopify-recruitment-backend/database"
	"github.com/Albert221/shopify-recruitment-backend/domain"
	"github.com/Albert221/shopify-recruitment-backend/resolver"
	"github.com/Albert221/shopify-recruitment-backend/stripe"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	db := createGorm()
	defer db.Close()

	// Resolver dependencies
	productsRepo := database.NewProductRepository(db)
	purchaseRepo := database.NewPurchaseRepository(db)
	stripeGate := stripe.PaymentGate{}

	rootResolver := resolver.NewRootResolver(productsRepo, purchaseRepo, stripeGate)

	schema := graphql.MustParseSchema(readSchema(), rootResolver)
	http.Handle("/api", &relay.Handler{Schema: schema})
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

	db.AutoMigrate(&domain.Product{}, &domain.ProductOrder{}, &domain.Purchase{})

	return db.Debug()
}
