package stripe

import "github.com/Albert221/shopify-recruitment-backend/domain/service"

type PaymentGate struct{}

func NewPaymentGate() *PaymentGate {
	return &PaymentGate{}
}

func (PaymentGate) Charge(cardDetails *service.CreditCardDetails, amount float64) error {
	// there is no logic for obvious reasons, pretend everything went nice

	return nil
}
