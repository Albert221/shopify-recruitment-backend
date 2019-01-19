package domain

type ProductRepository interface {
	Get(id string) *Product
	GetMany(ids []string) []*Product
	GetAll() []*Product
	GetAvailable() []*Product

	Save(product *Product) error
}
