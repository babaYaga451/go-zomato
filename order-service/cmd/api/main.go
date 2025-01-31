package main

import (
	"context"
	"time"

	"github.com/babaYaga451/go-zomato/common/log"
	"github.com/babaYaga451/go-zomato/common/uuid"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/config"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/inbound/http"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/inbound/http/handler"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/outbound/repository/mongodb"
	"github.com/babaYaga451/go-zomato/order-service/internal/application/service"
)

func main() {
	logger := log.NewZapLogger()
	uuidGenerator := uuid.NewRandomUUIDGeneratory()
	conf, err := config.New()
	if err != nil {
		logger.Fatalw("failed to load config", "error", err)
	}

	client, err := mongodb.New(conf.DB)
	dbName := conf.DB.DbName
	if err != nil {
		logger.Fatalw("failed to load config", "error", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)
	logger.Info("MongoDB connection established!")

	orderRepository := mongodb.NewOrderRepository(client, dbName)
	customerRepository := mongodb.NewCustomerRepository(client, dbName)
	restaurantRepository := mongodb.NewRestaurantRepository(client, dbName)

	orderService := service.NewOrderService(orderRepository, customerRepository, restaurantRepository, uuidGenerator, logger)
	orderHandler := handler.NewOrderCommandHandler(orderService, logger)

	router := http.NewRouterWithConfig(orderHandler, conf.HTTP, logger)
	mux := router.SetUpRouter()
	logger.Fatal(router.Run(mux))
}
