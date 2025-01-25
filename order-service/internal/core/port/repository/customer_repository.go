package repository

import (
	"context"

	"github.com/babaYaga451/go-zomato/order-service/internal/core/domain/entity"
)

type CustomerRepository interface {
	FindCustomer(ctx context.Context, customerID string) (*entity.Customer, error)
}
