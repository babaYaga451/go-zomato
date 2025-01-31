package http

import (
	"net/http"
	"time"

	"github.com/babaYaga451/go-zomato/common/log"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/config"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/inbound/http/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	orderCommandHandler *handler.OrderCommandHandler
	config              *config.HTTP
	logger              log.Logger
}

func NewRouterWithConfig(orderCommandHandler *handler.OrderCommandHandler, config *config.HTTP, logger log.Logger) *Router {
	return &Router{
		orderCommandHandler: orderCommandHandler,
		config:              config,
		logger:              logger,
	}
}

func (rtr *Router) SetUpRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Route("/orders", func(r chi.Router) {
			r.Post("/", rtr.orderCommandHandler.CreateOrderHandler)
			r.Get("/{trackingId}", rtr.orderCommandHandler.TrackOrderHandler)
		})
	})
	return r
}

func (rtr *Router) Run(mux http.Handler) error {
	server := http.Server{
		Handler:      mux,
		Addr:         rtr.config.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Minute,
	}

	rtr.logger.Infow("server started", "port", rtr.config.Port)
	return server.ListenAndServe()
}
