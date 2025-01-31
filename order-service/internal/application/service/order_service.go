package service

import (
	"context"

	"github.com/babaYaga451/go-zomato/common/log"
	"github.com/babaYaga451/go-zomato/common/uuid"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/inbound/http/dto"
	"github.com/babaYaga451/go-zomato/order-service/internal/application/mapper"
	"github.com/babaYaga451/go-zomato/order-service/internal/application/port"
	"github.com/babaYaga451/go-zomato/order-service/internal/domain/errors"
)

type OrderService struct {
	orderRepository      port.OrderRepository
	customerRepository   port.CustomerRepository
	restaurantRepository port.RestaurantRepository
	uuidGenerator        uuid.UUIDGenerator
	logger               log.Logger
}

func NewOrderService(orderRepository port.OrderRepository,
	customerRepository port.CustomerRepository,
	restaurantRepository port.RestaurantRepository,
	uuidGenerator uuid.UUIDGenerator,
	logger log.Logger) *OrderService {
	return &OrderService{
		orderRepository:      orderRepository,
		customerRepository:   customerRepository,
		restaurantRepository: restaurantRepository,
		uuidGenerator:        uuidGenerator,
		logger:               logger,
	}
}

func (os *OrderService) CreateOrder(ctx context.Context, cmd *dto.CreateOrderCommand) (*dto.CreateOrderResponse, error) {
	order := mapper.MapToDomainOrderEntity(cmd)
	_, err := os.customerRepository.FindCustomer(ctx, cmd.CustomerID)

	if err != nil {
		os.logger.Warn("Customer not found with id: ", cmd.CustomerID)
		return nil, err
	}

	productIds := make([]string, len(cmd.Items))
	for _, item := range cmd.Items {
		productIds = append(productIds, item.ProductId)
	}

	restaurant, err := os.restaurantRepository.FindRestaurantByProducts(ctx, cmd.RestaurantID, productIds)
	if err != nil {
		return nil, err
	}

	if !restaurant.IsActive() {
		return nil, errors.NewOrderDomainException("Restaurant with id: " + cmd.RestaurantID + " is not active")
	}

	err = order.CreateNewOrder(
		os.uuidGenerator.GenerateOrderID(),
		os.uuidGenerator.GenerateTrackingID(),
		order,
		restaurant)

	if err != nil {
		return nil, err
	}

	if err := os.orderRepository.SaveOrderAndInitiatePaymentTx(ctx, order); err != nil {
		os.logger.Error("Error in saving order!")
		return nil, errors.NewOrderDomainException("Error in saving order!")
	}

	os.logger.Info("Order created with id: " + order.GetID())
	return mapper.MapToOrderResponseDto(order), nil
}

func (os *OrderService) TrackOrder(ctx context.Context, trackingId string) (*dto.TrackOrderResponse, error) {
	order, err := os.orderRepository.FindByTrackingId(ctx, trackingId)

	if err != nil {
		os.logger.Warn("Order not found with tracking id: ", trackingId)
		return nil, errors.NewOrderNotFoundException("Order not found with tracking id: " + trackingId)
	}
	return mapper.MapToTrackingOrderResponseDto(order), nil
}
