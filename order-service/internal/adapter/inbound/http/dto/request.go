package dto

type CreateOrderCommand struct {
	CustomerID   string       `json:"customer_id"`
	RestaurantID string       `json:"restaurant_id"`
	Price        float64      `json:"price"`
	Items        []OrderItem  `json:"items"`
	Address      OrderAddress `json:"address"`
}

type OrderItem struct {
	ProductId string  `json:"product_id" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
	SubTotal  float64 `json:"sub_total" validate:"required"`
}

type OrderAddress struct {
	Street     string `json:"street" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
	City       string `json:"city" validate:"required"`
}
