package dto

type CreateOrderCommand struct {
	CustomerID   string       `json:"customer_id"`
	RestaurantID string       `json:"restaurant_id"`
	Price        float64      `json:"price"`
	Items        []OrderItem  `json:"items"`
	Address      OrderAddress `json:"address"`
}
