package main

import (
	"github.com/Albert221/shopify-recruitment-backend/model"
	"github.com/Albert221/shopify-recruitment-backend/resolver"
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

	db.AutoMigrate(&model.Product{}, &model.ProductOrder{}, &model.Purchase{})

	rootResolver := resolver.NewRootResolver(db)

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

	return db
}
