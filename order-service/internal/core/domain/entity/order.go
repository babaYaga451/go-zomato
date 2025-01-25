package entity

import (
	"fmt"

	"github.com/babaYaga451/go-zomato/common/common-domain/valueObject"
	"github.com/babaYaga451/go-zomato/order-service/internal/core/domain/exception"
	valueobject "github.com/babaYaga451/go-zomato/order-service/internal/core/domain/valueObject"
)

type Order struct {
	id              string
	customerID      string
	restaurantID    string
	deliveryAddress *valueobject.Address
	price           valueObject.Money
	items           []*OrderItem
	trackingID      string
	orderStatus     OrderStatus
	failureMessages []string
}

type OrderStatus string

func NewOrder(customerID string,
	restaurantID string,
	deliveryAddress *valueobject.Address,
	price valueObject.Money,
	items []*OrderItem) *Order {
	return &Order{
		customerID:      customerID,
		restaurantID:    restaurantID,
		deliveryAddress: deliveryAddress,
		price:           price,
		items:           items,
	}
}

func NewOrderWithTrackingID(
	orderID string,
	customerID string,
	restaurantID string,
	deliveryAddress *valueobject.Address,
	price valueObject.Money,
	items []*OrderItem,
	trackingID string,
	failureMessages []string,
	orderStatus OrderStatus) *Order {
	return &Order{
		id:              orderID,
		customerID:      customerID,
		restaurantID:    restaurantID,
		deliveryAddress: deliveryAddress,
		price:           price,
		items:           items,
		trackingID:      trackingID,
		failureMessages: failureMessages,
		orderStatus:     orderStatus,
	}
}

func (o *Order) GetID() string {
	return o.id
}

func (o *Order) GetCustomerID() string {
	return o.customerID
}

func (o *Order) GetRestaurantID() string {
	return o.restaurantID
}

func (o *Order) GetDeliveryAddress() *valueobject.Address {
	return o.deliveryAddress
}

func (o *Order) GetPrice() valueObject.Money {
	return o.price
}

func (o *Order) GetTrackingID() string {
	return o.trackingID
}

func (o *Order) GetOrderStatus() OrderStatus {
	return o.orderStatus
}

func (o *Order) GetOrderItems() []*OrderItem {
	return o.items
}

func (o *Order) GetFailureMessages() []string {
	return o.failureMessages
}

func (o *Order) SetID(id string) {
	o.id = id
}

func (o *Order) SetOrderStatus(status OrderStatus) {
	o.orderStatus = status
}

func (o *Order) SetTrackingID(trackingID string) {
	o.trackingID = trackingID
}

func (o *Order) ValidateOrder() error {
	if err := o.ensureValidIdAndStatus(); err != nil {
		return err
	}

	if err := o.validateTotalPrice(); err != nil {
		return err
	}

	if err := o.validateOrderItemsPrice(); err != nil {
		return err
	}

	return nil
}

func (o *Order) ensureValidIdAndStatus() error {
	if o.id != "" || o.orderStatus != "" {
		return exception.NewOrderDomainException("order is not in correct state for initialization")
	}

	return nil
}

func (o *Order) validateOrderItemsPrice() error {
	var orderItemTotalPrice = valueObject.NewMoney(0)
	for _, item := range o.items {
		if item.IsPriceValid() {
			orderItemTotalPrice = orderItemTotalPrice.Add(item.GetSubTotal())
		}
	}

	if !o.price.Equals(orderItemTotalPrice) {
		errMssg := fmt.Sprintf("total price %.2f is not equal to order items total price %.2f", o.price.GetAmount(), orderItemTotalPrice.GetAmount())
		return exception.NewOrderDomainException(errMssg)
	}
	return nil
}

func (o *Order) validateTotalPrice() error {
	if !o.price.IsGreaterThanZero() {
		return exception.NewOrderDomainException("total price must be greater than zero")
	}
	return nil
}

func (o *Order) IntializeOrderItems() {
	itemId := 1
	for _, item := range o.items {
		item.InitializeItem(o.id, itemId)
		itemId++
	}
}
