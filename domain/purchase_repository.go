package domain

type PurchaseRepository interface {
	Purchase(purchase *Purchase) error
}
