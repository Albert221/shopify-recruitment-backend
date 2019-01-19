package stripe

import "github.com/Albert221/shopify-recruitment-backend/domain/service"

type PaymentGate struct {}

func (PaymentGate) Charge(cardDetails *service.CreditCardDetails, amount float64) error {
	// This is an implementation that will NOT validate card details and charge the money

	return nil
}
