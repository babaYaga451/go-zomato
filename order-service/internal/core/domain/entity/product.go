package entity

import "github.com/babaYaga451/go-zomato/common/common-domain/valueObject"

type Product struct {
	id    string
	name  string
	price valueObject.Money
}

func NewProduct(id string) *Product {
	return &Product{
		id: id,
	}
}

func NewProductWithNameAndPrice(id string, name string, price valueObject.Money) *Product {
	return &Product{
		id:    id,
		name:  name,
		price: price,
	}
}

func (p *Product) GetID() string {
	return p.id
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetPrice() valueObject.Money {
	return p.price
}

func (p *Product) UpdateWithNameAndPrice(name string, price valueObject.Money) {
	p.name = name
	p.price = price
}
