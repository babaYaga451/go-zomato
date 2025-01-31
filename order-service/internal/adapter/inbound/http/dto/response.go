package dto

import "github.com/babaYaga451/go-zomato/order-service/internal/domain"

type CreateOrderResponse struct {
	TrackingID  string             `json:"tracking_id" validate:"required"`
	OrderStatus domain.OrderStatus `json:"order_status" validate:"required"`
	Message     string             `json:"message" validate:"required"`
}

type TrackOrderResponse struct {
	OrderTrackingId string             `json:"order_tracking_id" validate:"required"`
	OrderStatus     domain.OrderStatus `json:"order_status" validate:"required"`
	FailureMessages []string           `json:"failure_messages"`
}
