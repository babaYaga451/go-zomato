package domain

type Restaurant struct {
	id       string
	products []*Product
	active   bool
}

func NewRestaurant(id string, products []*Product, active bool) *Restaurant {
	return &Restaurant{
		id:       id,
		products: products,
		active:   active,
	}
}

func (r *Restaurant) GetID() string {
	return r.id
}

func (r *Restaurant) GetProducts() []*Product {
	return r.products
}

func (r *Restaurant) IsActive() bool {
	return r.active
}
