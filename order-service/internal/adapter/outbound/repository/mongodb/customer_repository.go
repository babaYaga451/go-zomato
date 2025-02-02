package mongodb

import (
	"context"

	"github.com/babaYaga451/go-zomato/order-service/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepositpry struct {
	client             *mongo.Client
	customerCollection *mongo.Collection
}

func NewCustomerRepository(client *mongo.Client, dbName string) *CustomerRepositpry {
	return &CustomerRepositpry{
		client:             client,
		customerCollection: client.Database(dbName).Collection("customers"),
	}
}

func (cr *CustomerRepositpry) FindCustomer(ctx context.Context, customerID string) (*domain.Customer, error) {
	return nil, nil
}
