package dto

type OrderAddress struct {
	Street     string `json:"street" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
	City       string `json:"city" validate:"required"`
}
