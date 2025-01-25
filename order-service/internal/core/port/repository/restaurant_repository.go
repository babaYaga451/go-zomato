package repository

import (
	"context"

	"github.com/babaYaga451/go-zomato/order-service/internal/core/domain/entity"
)

type RestaurantRepository interface {
	FindRestaurant(ctx context.Context, restaurantID string, productIds []string) (*entity.Restaurant, error)
}
