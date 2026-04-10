package dto

import "github.com/shopspring/decimal"

type UserResponse struct {
	UserID   int32   `json:"user_id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Location string  `json:"location"`
	Orders   []Order `json:"orders"`
}

type Order struct {
	UserID     int32           `json:"user_id"`
	Item       string          `json:"item"`
	Quantity   int32           `json:"quantity"`
	Price      decimal.Decimal `json:"price"`
	Discount   decimal.Decimal `json:"discount"`
	FinalPrice decimal.Decimal `json:"final_price"`
}
