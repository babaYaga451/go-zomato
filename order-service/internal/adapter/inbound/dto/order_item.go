package dto

type OrderItem struct {
	ProductId string  `json:"product_id" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
	SubTotal  float64 `json:"sub_total" validate:"required"`
}
