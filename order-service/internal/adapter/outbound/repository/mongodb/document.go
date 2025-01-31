package mongodb

import (
	"time"
)

type OrderDocument struct {
	OrderID         string                `bson:"_id,omitempty"`
	CustomerID      string                `bson:"customer_id"`
	RestaurantID    string                `bson:"restaurant_id"`
	OrderStatus     string                `bson:"order_status"`
	Price           float64               `bson:"price"`
	Items           []*OrderItemDocument  `bson:"items"`
	DeliveryAddress *OrderAddressDocument `bson:"delivery_address"`
	TrackingID      string                `bson:"tracking_id"`
	FailureMessages []string              `bson:"failure_messages"`
}

type OrderPaymnetOutboxPayload struct {
	Id                 string    `bson:"_id,omitempty"`
	CustomerID         string    `bson:"customer_id"`
	OrderID            string    `bson:"order_id"`
	Price              float64   `bson:"price"`
	CreatedAt          time.Time `bson:"created_at"`
	PaymentOrderStatus string    `bson:"payment_order_status"`
}

type RestaurantDocument struct {
	RestaurantID string             `bson:"_id,omitempty"`
	Products     []*ProductDocument `bson:"products"`
	Active       bool               `bson:"active"`
}

type OrderItemDocument struct {
	OrderID     string           `bson:"order_id"`
	OrderItemId int              `bson:"order_item_id"`
	Quantity    int              `bson:"quantity"`
	Subtotal    float64          `bson:"sub_total"`
	Price       float64          `bson:"price"`
	Product     *ProductDocument `bson:"product"`
}

type ProductDocument struct {
	ProductID string  `bson:"_id"`
	Name      string  `bson:"name"`
	Price     float64 `bson:"price"`
}

type OrderAddressDocument struct {
	Id         string `bson:"_id,omitempty"`
	PostalCode string `bson:"postal_code"`
	Street     string `bson:"street"`
	City       string `bson:"city"`
}
