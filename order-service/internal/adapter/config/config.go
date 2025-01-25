package config

import (
	"log"

	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/config/env"
	"github.com/joho/godotenv"
)

type (
	Container struct {
		App  *App
		DB   *DB
		HTTP *HTTP
	}

	App struct {
		Name string
		Env  string
	}

	DB struct {
		Uri    string
		DbName string
	}

	HTTP struct {
		Port string
	}
)

func New() (*Container, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	app := &App{
		Name: env.GetString("APP_NAME", "Order Service"),
		Env:  env.GetString("APP_ENV", "dev"),
	}

	db := &DB{
		Uri:    env.GetString("DB_URI", "mongodb://localhost:27017"),
		DbName: env.GetString("DB_NAME", "food-ordering"),
	}

	http := &HTTP{
		Port: env.GetString("HTTP_PORT", ":8080"),
	}

	return &Container{
		App:  app,
		DB:   db,
		HTTP: http,
	}, nil
}
