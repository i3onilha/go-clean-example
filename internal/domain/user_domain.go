package domain

import (
	"errors"

	"github.com/shopspring/decimal"
)

var ErrNotFound = errors.New("not found")

type Order struct {
	OrderID         int32           `json:"order_id"`
	UserID          int32           `json:"user_id"`
	Item            string          `json:"item"`
	Quantity        int32           `json:"quantity"`
	Price           decimal.Decimal `json:"price"`
	DiscountPercent decimal.Decimal `json:"discount_percent"`
}

// CalculateFinalPrice contains business logic that MUST live in the domain layer.
// Usecases delegate to this method — they never reimplement this logic.
func (d *Order) CalculateFinalPrice() decimal.Decimal {
	percent := d.DiscountPercent
	if percent.LessThan(decimal.Zero) {
		percent = decimal.Zero
	}
	if percent.GreaterThan(decimal.NewFromInt(100)) {
		percent = decimal.NewFromInt(100)
	}
	total := decimal.NewFromInt(int64(d.Quantity)).Mul(d.Price)
	discountFactor := decimal.NewFromInt(100).Sub(percent).Div(decimal.NewFromInt(100))
	return total.Mul(discountFactor)
}

type GetOrdersResponse struct {
	Orders []*Order `json:"orders"`
}

type GetUserInfoResponse struct {
	UserID   int32  `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Location string `json:"location"`
}
