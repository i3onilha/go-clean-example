package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go-clean-example/internal/domain"
	"go-clean-example/internal/repository"
	"go-clean-example/internal/repository/mysql/user"

	"github.com/shopspring/decimal"
)

type userRepository struct {
	query user.Queries
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{query: *user.New(db)}
}

func (r *userRepository) GetUserInfo(ctx context.Context, userID int32) (*domain.GetUserInfoResponse, error) {
	usr, err := r.query.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &domain.GetUserInfoResponse{
		UserID: usr.UserID,
		Name:   usr.Name,
		Email:  usr.Email,
		Location: func() string {
			if usr.Location.Valid {
				return usr.Location.String
			}
			return ""
		}(),
	}, nil
}

func (r *userRepository) GetOrdersByUserID(ctx context.Context, userID int32, limit, offset int) (*domain.GetOrdersResponse, error) {
	orders, err := r.query.GetOrdersByUserID(ctx, user.GetOrdersByUserIDParams{
		UserID: userID,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	ordersResponse := make([]*domain.Order, 0, len(orders))
	for _, o := range orders {
		var discountPercent decimal.Decimal
		switch v := o.DiscountPercent.(type) {
		case decimal.Decimal:
			discountPercent = v
		case []byte:
			discountPercent, _ = decimal.NewFromString(string(v))
		case float64:
			discountPercent = decimal.NewFromFloat(v)
		case nil:
			discountPercent = decimal.Zero
		default:
			return nil, fmt.Errorf("unexpected type for discount_percent: %T", v)
		}
		ordersResponse = append(ordersResponse, &domain.Order{
			UserID:          int32(o.UserID),
			OrderID:         int32(o.OrderID),
			Item:            o.Item,
			Quantity:        int32(o.Quantity),
			Price:           o.Price,
			DiscountPercent: discountPercent,
		})
	}
	return &domain.GetOrdersResponse{Orders: ordersResponse}, nil
}
