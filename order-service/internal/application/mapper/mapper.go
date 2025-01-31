package mapper

import (
	"github.com/babaYaga451/go-zomato/common/common-domain/valueObject"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/inbound/http/dto"
	"github.com/babaYaga451/go-zomato/order-service/internal/domain"
	valueobject "github.com/babaYaga451/go-zomato/order-service/internal/domain/valueObject"
)

func MapToDomainOrderEntity(orderCommandDto *dto.CreateOrderCommand) *domain.Order {
	return domain.NewOrderBuilder().
		WithCustomerId(orderCommandDto.CustomerID).
		WithRestaurantId(orderCommandDto.RestaurantID).
		WithDeliveryAddress(mapAddress(orderCommandDto.Address)).
		WithPrice(valueObject.NewMoney(orderCommandDto.Price)).
		WithItems(mapOrderItems(orderCommandDto.Items)).
		Build()
}

func MapToOrderResponseDto(order *domain.Order) *dto.CreateOrderResponse {
	return &dto.CreateOrderResponse{
		TrackingID:  order.GetTrackingID(),
		OrderStatus: order.GetOrderStatus(),
		Message:     "Order created successfully",
	}
}

func mapAddress(orderAddress dto.OrderAddress) *valueobject.Address {
	return valueobject.NewAddress(
		orderAddress.PostalCode,
		orderAddress.Street,
		orderAddress.City)
}

func mapOrderItems(orderItems []dto.OrderItem) []*domain.OrderItem {
	items := make([]*domain.OrderItem, len(orderItems))
	for i, item := range orderItems {
		product := domain.NewProduct(item.ProductId)
		price := valueObject.NewMoney(item.Price)
		subTotal := valueObject.NewMoney(item.SubTotal)

		items[i] = domain.NewOrderItem(product, item.Quantity, price, subTotal)
	}
	return items
}

func MapToTrackingOrderResponseDto(order *domain.Order) *dto.TrackOrderResponse {
	return &dto.TrackOrderResponse{
		OrderTrackingId: order.GetTrackingID(),
		OrderStatus:     order.GetOrderStatus(),
		FailureMessages: order.GetFailureMessages(),
	}
}
