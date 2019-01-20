package service

type CreditCardDetails struct {
	Holder  string
	Number  string
	Expires string
	CVV     string
}

type PaymentGate interface {
	Charge(cardDetails *CreditCardDetails, amount float64) error
}
