package domain

import (
	"github.com/babaYaga451/go-zomato/common/common-domain/valueObject"
	valueobject "github.com/babaYaga451/go-zomato/order-service/internal/domain/valueObject"
)

type OrderBuilder struct {
	id              string
	customerID      string
	restaurantID    string
	deliveryAddress *valueobject.Address
	price           valueObject.Money
	items           []*OrderItem
	trackingID      string
	orderStatus     OrderStatus
}

func NewOrderBuilder() *OrderBuilder {
	return &OrderBuilder{}
}

func (ob *OrderBuilder) WithId(id string) *OrderBuilder {
	ob.id = id
	return ob
}

func (ob *OrderBuilder) WithCustomerId(customerId string) *OrderBuilder {
	ob.customerID = customerId
	return ob
}

func (ob *OrderBuilder) WithRestaurantId(restaurantId string) *OrderBuilder {
	ob.restaurantID = restaurantId
	return ob
}

func (ob *OrderBuilder) WithDeliveryAddress(deliveryAddress *valueobject.Address) *OrderBuilder {
	ob.deliveryAddress = deliveryAddress
	return ob
}

func (ob *OrderBuilder) WithPrice(price valueObject.Money) *OrderBuilder {
	ob.price = price
	return ob
}

func (ob *OrderBuilder) WithItems(items []*OrderItem) *OrderBuilder {
	ob.items = items
	return ob
}

func (ob *OrderBuilder) WithTrackingId(trackingId string) *OrderBuilder {
	ob.trackingID = trackingId
	return ob
}

func (ob *OrderBuilder) WithOrderStatus(orderStatus OrderStatus) *OrderBuilder {
	ob.orderStatus = orderStatus
	return ob
}

func (ob *OrderBuilder) Build() *Order {
	return &Order{
		id:              ob.id,
		customerID:      ob.customerID,
		restaurantID:    ob.restaurantID,
		deliveryAddress: ob.deliveryAddress,
		price:           ob.price,
		items:           ob.items,
		trackingID:      ob.trackingID,
		orderStatus:     ob.orderStatus,
	}
}
