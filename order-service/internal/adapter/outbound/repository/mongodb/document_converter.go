package mongodb

import (
	"time"

	"github.com/babaYaga451/go-zomato/common/common-domain/valueObject"
	"github.com/babaYaga451/go-zomato/order-service/internal/domain"
	valueobject "github.com/babaYaga451/go-zomato/order-service/internal/domain/valueObject"
	"github.com/google/uuid"
)

func ToOrderDocument(order *domain.Order) *OrderDocument {
	return &OrderDocument{
		OrderID:         order.GetID(),
		CustomerID:      order.GetCustomerID(),
		RestaurantID:    order.GetRestaurantID(),
		OrderStatus:     string(order.GetOrderStatus()),
		Price:           order.GetPrice().GetAmount(),
		Items:           toOrderItemsDocument(order.GetOrderItems()),
		DeliveryAddress: toOrderAddressDocument(order.GetDeliveryAddress()),
		TrackingID:      order.GetTrackingID(),
		FailureMessages: order.GetFailureMessages(),
	}
}

func ToOrderDomainEntity(orderDocument *OrderDocument) *domain.Order {
	return domain.NewOrderBuilder().
		WithId(orderDocument.OrderID).
		WithCustomerId(orderDocument.CustomerID).
		WithRestaurantId(orderDocument.RestaurantID).
		WithDeliveryAddress(toAddress(orderDocument.DeliveryAddress)).
		WithPrice(valueObject.NewMoney(orderDocument.Price)).
		WithItems(toOrderItems(orderDocument.Items)).
		WithTrackingId(orderDocument.TrackingID).
		WithOrderStatus(domain.OrderStatus(orderDocument.OrderStatus)).
		Build()
}

func ToOrderPaymnetOutboxPayload(payload, orderID string) *OrderPaymnetOutboxPayload {
	return &OrderPaymnetOutboxPayload{
		Id:                 uuid.New().String(),
		OrderID:            orderID,
		Payload:            payload,
		OutboxStatus:       "STARTED",
		ProcessingAttempts: 0,
		CreatedAt:          time.Now(),
	}
}

func ToRestaurantDomainEntity(restaurantDocument *RestaurantDocument) *domain.Restaurant {
	return domain.NewRestaurant(restaurantDocument.RestaurantID,
		toProducts(restaurantDocument.Products),
		restaurantDocument.Active)
}

func toProducts(productDocument []*ProductDocument) []*domain.Product {
	products := make([]*domain.Product, len(productDocument))
	for i := range productDocument {
		product := productDocument[i]
		products[i] = domain.NewProductWithNameAndPrice(product.ProductID, product.Name, valueObject.NewMoney(product.Price))
	}
	return products
}

func toOrderAddressDocument(address *valueobject.Address) *OrderAddressDocument {
	return &OrderAddressDocument{
		PostalCode: address.GetPostalCode(),
		Street:     address.GetStreet(),
		City:       address.GetCity(),
	}
}

func toOrderItemsDocument(orderItem []*domain.OrderItem) []*OrderItemDocument {
	items := make([]*OrderItemDocument, len(orderItem))
	for i := range orderItem {
		item := orderItem[i]
		items[i] = &OrderItemDocument{
			OrderID:     item.GetOrderID(),
			OrderItemId: item.GetOrderItemID(),
			Quantity:    item.GetQuantity(),
			Subtotal:    item.GetSubTotal().GetAmount(),
			Price:       item.GetPrice().GetAmount(),
			Product:     toProductDocument(item.GetProduct()),
		}
	}
	return items
}

func toProductDocument(product *domain.Product) *ProductDocument {
	return &ProductDocument{
		ProductID: product.GetID(),
		Name:      product.GetName(),
		Price:     product.GetPrice().GetAmount(),
	}
}

func toOrderItems(orderItemDocument []*OrderItemDocument) []*domain.OrderItem {
	orderItems := make([]*domain.OrderItem, len(orderItemDocument))
	for i := range orderItemDocument {
		item := orderItemDocument[i]
		productId := item.Product.ProductID
		productName := item.Product.Name
		productPrice := valueObject.NewMoney(item.Product.Price)
		product := domain.NewProductWithNameAndPrice(productId, productName, productPrice)
		orderItemPrice := valueObject.NewMoney(item.Price)
		orderItemSubtotal := valueObject.NewMoney(item.Subtotal)
		orderItems[i] = domain.NewOrderItem(
			product,
			item.Quantity,
			orderItemPrice,
			orderItemSubtotal,
		)
	}
	return orderItems
}

func toAddress(orderAddressDocument *OrderAddressDocument) *valueobject.Address {
	return valueobject.NewAddress(
		orderAddressDocument.PostalCode,
		orderAddressDocument.Street,
		orderAddressDocument.City)
}
