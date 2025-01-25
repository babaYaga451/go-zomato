package dto

import "github.com/babaYaga451/go-zomato/order-service/internal/core/domain/entity"

type TrackOrderResponse struct {
	OrderTrackingId string             `json:"order_tracking_id" validate:"required"`
	OrderStatus     entity.OrderStatus `json:"order_status" validate:"required"`
	FailureMessages []string           `json:"failure_messages"`
}
