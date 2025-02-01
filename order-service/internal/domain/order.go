package domain

import (
	"fmt"
	"time"

	"github.com/babaYaga451/go-zomato/common/common-domain/valueObject"
	"github.com/babaYaga451/go-zomato/order-service/internal/domain/errors"
	"github.com/babaYaga451/go-zomato/order-service/internal/domain/event"
	valueobject "github.com/babaYaga451/go-zomato/order-service/internal/domain/valueObject"
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

func (o *Order) CreateNewOrder(orderId, trackingId string, restaurant *Restaurant) (*Order, *event.OrderPaymentEvent, error) {
	if err := o.ensureValidIdAndStatus(); err != nil {
		return nil, nil, err
	}

	if err := o.validateTotalPrice(); err != nil {
		return nil, nil, err
	}

	if err := o.validateRestaurant(restaurant); err != nil {
		return nil, nil, err
	}

	newOrder := &Order{
		id:              orderId,
		customerID:      o.customerID,
		restaurantID:    o.restaurantID,
		deliveryAddress: o.deliveryAddress,
		price:           o.price,
		items:           o.items,
		trackingID:      trackingId,
		orderStatus:     "PENDING",
		failureMessages: nil,
	}
	newOrder.intializeOrderItems()

	return newOrder,
		&event.OrderPaymentEvent{
			OrderID:       newOrder.id,
			CustomerID:    newOrder.customerID,
			Price:         newOrder.price.GetAmount(),
			PaymentStatus: "PENDING",
			CreatedAt:     time.Now(),
		}, nil
}

func (o *Order) ensureValidIdAndStatus() error {
	if o.id != "" || o.orderStatus != "" {
		return errors.NewOrderDomainException("order is not in correct state for initialization")
	}
	return nil
}

func (o *Order) ValidateOrderItemsPrice() error {
	var total float64 = 0
	for _, item := range o.items {
		if item.IsPriceValid() {
			total += item.GetSubTotal().GetAmount()
		}
	}

	if o.price.GetAmount() != total {
		errMssg := fmt.Sprintf("total price %.2f is not equal to order items total price %.2f", o.price.GetAmount(), total)
		return errors.NewOrderDomainException(errMssg)
	}
	return nil
}

func (o *Order) validateTotalPrice() error {
	if !o.price.IsGreaterThanZero() {
		return errors.NewOrderDomainException("total price must be greater than zero")
	}
	return nil
}

func (o *Order) intializeOrderItems() {
	itemId := 1
	for _, item := range o.items {
		item.InitializeItem(o.id, itemId)
		itemId++
	}
}

func (o *Order) validateRestaurant(restaurant *Restaurant) error {
	if !restaurant.IsActive() {
		return errors.NewOrderDomainException("restaurant is not active")
	}
	return nil
}

func (o *Order) SetOrderProductInformation(restaurant *Restaurant) {
	productMap := make(map[string]*Product)
	for _, product := range restaurant.GetProducts() {
		productMap[product.GetID()] = product
	}

	for _, item := range o.items {
		productId := item.GetProductId()
		product, exists := productMap[productId]
		if exists {
			item.SetProductInformation(product.GetName(), product.GetPrice())
		}
	}
}
