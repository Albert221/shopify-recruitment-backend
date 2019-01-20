package service

// CreditCardDetails contains credit card info.
type CreditCardDetails struct {
	Holder  string
	Number  string
	Expires string
	CVV     string
}

// PaymentGate is a service that can charge credit cards.
type PaymentGate interface {
	Charge(cardDetails *CreditCardDetails, amount float64) error
}
