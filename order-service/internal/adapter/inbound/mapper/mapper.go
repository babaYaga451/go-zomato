package mapper

import (
	"github.com/babaYaga451/go-zomato/common/common-domain/valueObject"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/inbound/dto"
	"github.com/babaYaga451/go-zomato/order-service/internal/core/domain/entity"
	valueobject "github.com/babaYaga451/go-zomato/order-service/internal/core/domain/valueObject"
	"github.com/google/uuid"
)

func MapToDomainOrderEntity(orderCommandDto *dto.CreateOrderCommand) *entity.Order {
	return entity.NewOrder(
		orderCommandDto.CustomerID,
		orderCommandDto.RestaurantID,
		mapAddress(orderCommandDto.Address),
		valueObject.NewMoney(orderCommandDto.Price),
		mapOrderItems(orderCommandDto.Items))
}

func MapToOrderResponseDto(order *entity.Order) *dto.CreateOrderResponse {
	return &dto.CreateOrderResponse{
		TrackingID:  order.GetTrackingID(),
		OrderStatus: order.GetOrderStatus(),
		Message:     "Order created successfully",
	}
}

func mapAddress(orderAddress dto.OrderAddress) *valueobject.Address {
	return valueobject.NewAddress(
		uuid.New().String(),
		orderAddress.PostalCode,
		orderAddress.Street,
		orderAddress.City)
}

func mapOrderItems(orderItem []dto.OrderItem) []*entity.OrderItem {
	var orderItems []*entity.OrderItem
	for _, item := range orderItem {
		orderItems = append(orderItems,
			entity.NewOrderItem(
				entity.NewProduct(item.ProductId),
				item.Quantity,
				valueObject.NewMoney(item.Price),
				valueObject.NewMoney(item.SubTotal)))
	}
	return orderItems
}

func MapToTrackingOrderResponseDto(order *entity.Order) *dto.TrackOrderResponse {
	return &dto.TrackOrderResponse{
		OrderTrackingId: order.GetTrackingID(),
		OrderStatus:     order.GetOrderStatus(),
		FailureMessages: order.GetFailureMessages(),
	}
}
