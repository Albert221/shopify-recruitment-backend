package service

type CreditCardDetails struct {
	Holder string
	Number string
	Expires int
	CVV int
}

type PaymentGate interface {
	Charge(cardDetails *CreditCardDetails, amount float64) error
}
