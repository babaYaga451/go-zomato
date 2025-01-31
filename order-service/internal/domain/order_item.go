package domain

import "github.com/babaYaga451/go-zomato/common/common-domain/valueObject"

type OrderItem struct {
	orderId     string
	orderItemId int
	product     *Product
	quantity    int
	price       valueObject.Money
	subTotal    valueObject.Money
}

func NewOrderItem(product *Product, quantity int, price valueObject.Money, subTotal valueObject.Money) *OrderItem {
	return &OrderItem{
		product:  product,
		quantity: quantity,
		price:    price,
		subTotal: subTotal,
	}
}

func (oi *OrderItem) GetOrderItemID() int {
	return oi.orderItemId
}

func (oi *OrderItem) GetOrderID() string {
	return oi.orderId
}

func (oi *OrderItem) GetQuantity() int {
	return oi.quantity
}

func (oi *OrderItem) GetPrice() valueObject.Money {
	return oi.price
}

func (oi *OrderItem) GetProduct() *Product {
	return oi.product
}

func (oi *OrderItem) GetProductId() string {
	return oi.product.id
}

func (oi *OrderItem) SetProductInformation(name string, price valueObject.Money) {
	oi.product.UpdateWithNameAndPrice(name, price)
}

func (oi *OrderItem) IsPriceValid() bool {
	itemsPrice := oi.price.Multiply(oi.quantity)

	return oi.price.IsGreaterThanZero() &&
		oi.price.Equals(oi.product.price) &&
		itemsPrice.Equals(oi.subTotal)
}

func (oi *OrderItem) GetSubTotal() valueObject.Money {
	return oi.subTotal
}

func (oi *OrderItem) InitializeItem(orderId string, orderItemId int) {
	oi.orderId = orderId
	oi.orderItemId = orderItemId
}
