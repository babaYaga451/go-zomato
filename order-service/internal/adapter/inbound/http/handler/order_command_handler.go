package handler

import (
	"net/http"

	"github.com/babaYaga451/go-zomato/common/json"
	"github.com/babaYaga451/go-zomato/common/log"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/inbound/http/dto"
	"github.com/babaYaga451/go-zomato/order-service/internal/application/port"
	"github.com/go-chi/chi/v5"
)

type OrderCommandHandler struct {
	orderService port.OrderService
	logger       log.Logger
}

func NewOrderCommandHandler(orderService port.OrderService, logger log.Logger) *OrderCommandHandler {
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
	response, err := h.orderService.CreateOrder(ctx, &createOrderCommand)
	if err != nil {
		h.internalServerError(w, r, err)
		return
	}
	json.WriteJSON(w, http.StatusCreated, response)
}

func (h *OrderCommandHandler) TrackOrderHandler(w http.ResponseWriter, r *http.Request) {
	trackingId := chi.URLParam(r, "trackingId")
	ctx := r.Context()
	response, err := h.orderService.TrackOrder(ctx, trackingId)
	if err != nil {
		h.notFoundError(w, r, err)
		return
	}
	json.WriteJSON(w, http.StatusOK, response)
}
