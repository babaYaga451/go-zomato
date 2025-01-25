package main

import (
	"context"
	"time"

	"github.com/babaYaga451/go-zomato/common/log"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/config"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/inbound/handler/http"
	mongodb "github.com/babaYaga451/go-zomato/order-service/internal/adapter/outbound/mongoDb"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/outbound/mongoDb/repository"
	"github.com/babaYaga451/go-zomato/order-service/internal/core/service"
)

func main() {
	logger := log.NewZapLogger()
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

	orderRepository := repository.NewOrderRepository(client, dbName)
	customerRepository := repository.NewCustomerRepository(client, dbName)
	restaurantRepository := repository.NewRestaurantRepository(client, dbName)

	orderService := service.NewOrderService(orderRepository, customerRepository, restaurantRepository, logger)
	orderHandler := http.NewOrderCommandHandler(orderService, logger)

	router := http.NewRouterWithConfig(orderHandler, conf.HTTP, logger)
	mux := router.SetUpRouter()
	logger.Fatal(router.Run(mux))
}
