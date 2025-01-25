package dto

import (
	"github.com/babaYaga451/go-zomato/order-service/internal/core/domain/entity"
)

type CreateOrderResponse struct {
	TrackingID  string             `json:"tracking_id" validate:"required"`
	OrderStatus entity.OrderStatus `json:"order_status" validate:"required"`
	Message     string             `json:"message" validate:"required"`
}
