package service

import (
	"context"

	"github.com/babaYaga451/go-zomato/common/log"
	"github.com/babaYaga451/go-zomato/order-service/internal/core/domain/entity"
	"github.com/babaYaga451/go-zomato/order-service/internal/core/domain/exception"
	"github.com/babaYaga451/go-zomato/order-service/internal/core/port/repository"
	"github.com/google/uuid"
)

type OrderServiceImpl struct {
	orderRepository      repository.OrderRepository
	customerRepository   repository.CustomerRepository
	restaurantRepository repository.RestaurantRepository
	logger               log.Logger
}

func NewOrderService(orderRepository repository.OrderRepository,
	customerRepository repository.CustomerRepository,
	restaurantRepository repository.RestaurantRepository,
	logger log.Logger) *OrderServiceImpl {
	return &OrderServiceImpl{
		orderRepository:      orderRepository,
		customerRepository:   customerRepository,
		restaurantRepository: restaurantRepository,
		logger:               logger,
	}
}

func (os *OrderServiceImpl) ValidateAndInitiateOrder(ctx context.Context, order *entity.Order) error {
	if err := os.validateCustomer(ctx, order.GetCustomerID()); err != nil {
		return err
	}

	restaurant, err := os.validateRestaurant(ctx, order)
	if err != nil {
		return err
	}

	if err := order.ValidateOrder(); err != nil {
		return err
	}

	os.setOrderProductInformation(order, restaurant)
	order.SetID(uuid.New().String())
	order.SetOrderStatus(entity.OrderStatus("PENDING"))
	order.SetTrackingID(uuid.New().String())
	order.IntializeOrderItems()

	if err := os.orderRepository.SaveOrderAndInitiatePaymentTx(ctx, order); err != nil {
		os.logger.Error("Error in saving order!")
		return exception.NewOrderDomainException("Error in saving order!")
	}

	os.logger.Info("Order created with id: " + order.GetID())
	return nil
}

func (os *OrderServiceImpl) TrackOrder(ctx context.Context, trackingId string) (*entity.Order, error) {
	order, err := os.orderRepository.FindByTrackingId(ctx, trackingId)

	if err != nil {
		os.logger.Warn("Order not found with tracking id: ", trackingId)
		return nil, exception.NewOrderNotFoundException("Order not found with tracking id: " + trackingId)
	}
	return order, nil
}

func (os *OrderServiceImpl) validateCustomer(ctx context.Context, customerID string) error {
	_, err := os.customerRepository.FindCustomer(ctx, customerID)

	if err != nil {
		os.logger.Warn("Customer not found with id: ", customerID)
		return exception.NewOrderDomainException("Customer not found with id: " + customerID)
	}

	return nil
}

func (os *OrderServiceImpl) validateRestaurant(ctx context.Context, order *entity.Order) (*entity.Restaurant, error) {
	var productIds []string
	restaurantId := order.GetRestaurantID()

	for _, item := range order.GetOrderItems() {
		productIds = append(productIds, item.GetProductId())
	}

	restaurant, err := os.restaurantRepository.FindRestaurant(ctx, restaurantId, productIds)
	if err != nil {
		os.logger.Warn("Restaurant not found with id: ", restaurantId)
		return nil, exception.NewOrderDomainException("Restaurant not found with id: " + restaurantId)
	}

	if !restaurant.IsActive() {
		return nil, exception.NewOrderDomainException("Restaurant with id: " + restaurantId + " is not active")
	}

	return restaurant, nil
}

func (os *OrderServiceImpl) setOrderProductInformation(order *entity.Order, restaurant *entity.Restaurant) {
	productMap := make(map[string]*entity.Product)
	for _, product := range restaurant.GetProducts() {
		productMap[product.GetID()] = product
	}

	orderItems := order.GetOrderItems()
	for i := range orderItems {
		productId := order.GetOrderItems()[i].GetProductId()
		product, exists := productMap[productId]
		if exists {
			order.GetOrderItems()[i].SetProductInformation(product.GetName(), product.GetPrice())
		} else {
			os.logger.Warnw("Product not found for id: ", productId)
		}
	}
}
