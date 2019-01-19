package domain

type CartRepository interface {
	Get(id string) *Cart
	Save(cart *Cart) error
}
