package resolver

import (
	"fmt"
	"github.com/Albert221/shopify-recruitment-backend/model"
)

type Money struct {
	model.Money
}

func (Money) ImplementsGraphQLType(name string) bool {
	return name == "Money"
}

func (m *Money) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case float64:
		m.Money = model.Money(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

