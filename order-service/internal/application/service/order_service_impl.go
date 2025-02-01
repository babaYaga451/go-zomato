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

type OrderServiceImpl struct {
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
	logger log.Logger) port.OrderService {
	return &OrderServiceImpl{
		orderRepository:      orderRepository,
		customerRepository:   customerRepository,
		restaurantRepository: restaurantRepository,
		uuidGenerator:        uuidGenerator,
		logger:               logger,
	}
}

func (os *OrderServiceImpl) CreateOrder(ctx context.Context, cmd *dto.CreateOrderCommand) (*dto.CreateOrderResponse, error) {
	_, err := os.customerRepository.FindCustomer(ctx, cmd.CustomerID)

	if err != nil {
		os.logger.Warn("Customer not found with id: ", cmd.CustomerID)
		return nil, err
	}

	productIds := make([]string, len(cmd.Items))
	for i, item := range cmd.Items {
		productIds[i] = item.ProductId
	}
	restaurant, err := os.restaurantRepository.FindRestaurantByProducts(ctx, cmd.RestaurantID, productIds)
	if err != nil {
		os.logger.Warn("Restaurant not found with id: ", cmd.RestaurantID)
		return nil, err
	}

	order := mapper.MapToDomainOrderEntity(cmd)
	order.SetOrderProductInformation(restaurant)
	if validationErr := order.ValidateOrderItemsPrice(); validationErr != nil {
		return nil, validationErr
	}

	newOrder, orderPaymentEvent, err := order.CreateNewOrder(
		os.uuidGenerator.GenerateOrderID(),
		os.uuidGenerator.GenerateTrackingID(),
		restaurant)

	if err != nil {
		return nil, err
	}

	err = os.orderRepository.SaveOrderAndInitiatePaymentTx(
		ctx,
		newOrder,
		orderPaymentEvent)

	if err != nil {
		os.logger.Error("Error in saving order!")
		return nil, errors.NewOrderDomainException("Error in saving order!")
	}

	os.logger.Info("Order created with id: " + order.GetID())
	return mapper.MapToOrderResponseDto(newOrder), nil
}

func (os *OrderServiceImpl) TrackOrder(ctx context.Context, trackingId string) (*dto.TrackOrderResponse, error) {
	order, err := os.orderRepository.FindByTrackingId(ctx, trackingId)

	if err != nil {
		os.logger.Warn("Order not found with tracking id: ", trackingId)
		return nil, errors.NewOrderNotFoundException("Order not found with tracking id: " + trackingId)
	}
	return mapper.MapToTrackingOrderResponseDto(order), nil
}
