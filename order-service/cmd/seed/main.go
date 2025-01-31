package main

import (
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/config"
	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/outbound/repository/mongodb"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		panic(err)
	}

	client, err := mongodb.New(cfg.DB)
	if err != nil {
		panic(err)
	}

	err = mongodb.SeedDatabase(client, cfg.DB.DbName)
	if err != nil {
		panic(err)
	}
}
