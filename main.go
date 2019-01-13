package main

import (
	"github.com/Albert221/shopify-recruitment-backend/resolver"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	rootResolver := &resolver.RootResolver{}

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
