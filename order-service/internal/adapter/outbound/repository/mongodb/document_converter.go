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

func ToOrderPaymnetOutboxPayload(order *domain.Order) *OrderPaymnetOutboxPayload {
	return &OrderPaymnetOutboxPayload{
		Id:                 uuid.New().String(),
		CustomerID:         order.GetCustomerID(),
		OrderID:            order.GetID(),
		Price:              order.GetPrice().GetAmount(),
		CreatedAt:          time.Now(),
		PaymentOrderStatus: string(order.GetOrderStatus()),
	}
}

func ToRestaurantDomainEntity(restaurantDocument *RestaurantDocument) *domain.Restaurant {
	return domain.NewRestaurant(restaurantDocument.RestaurantID,
		toProducts(restaurantDocument.Products),
		restaurantDocument.Active)
}

func toProducts(productDocument []*ProductDocument) []*domain.Product {
	var products []*domain.Product
	for _, product := range productDocument {
		products = append(products, domain.NewProductWithNameAndPrice(product.ProductID, product.Name, valueObject.NewMoney(product.Price)))
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
	var items []*OrderItemDocument
	for _, item := range orderItem {
		items = append(items, &OrderItemDocument{
			OrderID:     item.GetOrderID(),
			OrderItemId: item.GetOrderItemID(),
			Quantity:    item.GetQuantity(),
			Subtotal:    item.GetSubTotal().GetAmount(),
			Price:       item.GetPrice().GetAmount(),
			Product:     toProductDocument(item.GetProduct()),
		})
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
	var orderItems []*domain.OrderItem
	for _, item := range orderItemDocument {
		orderItems = append(orderItems, domain.NewOrderItem(
			domain.NewProductWithNameAndPrice(item.Product.ProductID, item.Product.Name, valueObject.NewMoney(item.Product.Price)),
			item.Quantity,
			valueObject.NewMoney(item.Price),
			valueObject.NewMoney(item.Subtotal),
		))
	}
	return orderItems
}

func toAddress(orderAddressDocument *OrderAddressDocument) *valueobject.Address {
	return valueobject.NewAddress(
		orderAddressDocument.PostalCode,
		orderAddressDocument.Street,
		orderAddressDocument.City)
}
