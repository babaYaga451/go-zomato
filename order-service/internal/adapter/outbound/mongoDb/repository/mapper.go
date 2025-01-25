package repository

import (
	"time"

	"github.com/babaYaga451/go-zomato/common/common-domain/valueObject"
	"github.com/babaYaga451/go-zomato/order-service/internal/core/domain/entity"
	valueobject "github.com/babaYaga451/go-zomato/order-service/internal/core/domain/valueObject"
)

func ToOrderDocument(order *entity.Order) *OrderDocument {
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

func ToOrderDomainEntity(orderDocument *OrderDocument) *entity.Order {
	return entity.NewOrderWithTrackingID(
		orderDocument.OrderID,
		orderDocument.CustomerID,
		orderDocument.RestaurantID,
		toAddress(orderDocument.DeliveryAddress),
		valueObject.NewMoney(orderDocument.Price),
		toOrderItems(orderDocument.Items),
		orderDocument.TrackingID,
		orderDocument.FailureMessages,
		entity.OrderStatus(orderDocument.OrderStatus))
}

func ToOrderPaymentOutboxDocument(order *entity.Order) *OrderPaymentOutboxDocument {
	return &OrderPaymentOutboxDocument{
		CustomerID:         order.GetCustomerID(),
		OrderID:            order.GetID(),
		price:              order.GetPrice().GetAmount(),
		createdAt:          time.Now(),
		paymentOrderStatus: string(order.GetOrderStatus()),
	}
}

func ToRestaurantDomainEntity(restaurantDocument *RestaurantDocument) *entity.Restaurant {
	return entity.NewRestaurant(restaurantDocument.RestaurantID,
		toProducts(restaurantDocument.Products),
		restaurantDocument.Active)
}

func toProducts(productDocument []*ProductDocument) []*entity.Product {
	var products []*entity.Product
	for _, product := range productDocument {
		products = append(products, entity.NewProductWithNameAndPrice(product.ProductID, product.Name, valueObject.NewMoney(product.Price)))
	}
	return products
}

func toOrderAddressDocument(address *valueobject.Address) *OrderAddressDocument {
	return &OrderAddressDocument{
		Id:         address.GetID(),
		PostalCode: address.GetPostalCode(),
		Street:     address.GetStreet(),
		City:       address.GetCity(),
	}
}

func toOrderItemsDocument(orderItem []*entity.OrderItem) []*OrderItemDocument {
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

func toProductDocument(product *entity.Product) *ProductDocument {
	return &ProductDocument{
		ProductID: product.GetID(),
		Name:      product.GetName(),
		Price:     product.GetPrice().GetAmount(),
	}
}

func toOrderItems(orderItemDocument []*OrderItemDocument) []*entity.OrderItem {
	var orderItems []*entity.OrderItem
	for _, item := range orderItemDocument {
		orderItems = append(orderItems, entity.NewOrderItem(
			entity.NewProductWithNameAndPrice(item.Product.ProductID, item.Product.Name, valueObject.NewMoney(item.Product.Price)),
			item.Quantity,
			valueObject.NewMoney(item.Price),
			valueObject.NewMoney(item.Subtotal),
		))
	}
	return orderItems
}

func toAddress(orderAddressDocument *OrderAddressDocument) *valueobject.Address {
	return valueobject.NewAddress(
		orderAddressDocument.Id,
		orderAddressDocument.PostalCode,
		orderAddressDocument.Street,
		orderAddressDocument.City)
}
