package repository

import (
	"context"

	"go-clean-example/internal/domain"
)

type UserRepository interface {
	GetUserInfo(ctx context.Context, userID int32) (*domain.GetUserInfoResponse, error)
	GetOrdersByUserID(ctx context.Context, userID int32, limit, offset int) (*domain.GetOrdersResponse, error)
}
