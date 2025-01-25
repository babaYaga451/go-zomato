package http

import (
	"net/http"

	"github.com/babaYaga451/go-zomato/common/json"
	"github.com/babaYaga451/go-zomato/common/log"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/inbound/dto"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/inbound/mapper"
	"github.com/babaYaga451/go-zomato/order-service/internal/core/port/service"
	"github.com/go-chi/chi/v5"
)

type OrderCommandHandler struct {
	orderService service.OrderService
	logger       log.Logger
}

func NewOrderCommandHandler(orderService service.OrderService, logger log.Logger) *OrderCommandHandler {
	return &OrderCommandHandler{
		orderService: orderService,
		logger:       logger,
	}
}

func (h *OrderCommandHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var createOrderCommand dto.CreateOrderCommand

	if err := json.ReadJSON(w, r, &createOrderCommand); err != nil {
		h.badRequestError(w, r, err)
		return
	}

	if err := json.Validate.Struct(createOrderCommand); err != nil {
		h.badRequestError(w, r, err)
		return
	}

	ctx := r.Context()
	order := mapper.MapToDomainOrderEntity(&createOrderCommand)

	err := h.orderService.ValidateAndInitiateOrder(ctx, order)
	if err != nil {
		h.internalServerError(w, r, err)
		return
	}

	response := mapper.MapToOrderResponseDto(order)
	json.WriteJSON(w, http.StatusOK, response)
}

func (h *OrderCommandHandler) TrackOrderHandler(w http.ResponseWriter, r *http.Request) {
	trackingId := chi.URLParam(r, "trackingId")

	ctx := r.Context()
	order, err := h.orderService.TrackOrder(ctx, trackingId)

	response := mapper.MapToTrackingOrderResponseDto(order)
	if err != nil {
		h.notFoundError(w, r, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, response)
}
