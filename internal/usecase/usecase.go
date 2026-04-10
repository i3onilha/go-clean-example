package usecase

import (
	"context"

	"go-clean-example/internal/dto"
)

type UserUsecase interface {
	GetUserWithOrders(ctx context.Context, id int32, limit, offset int) (*dto.UserResponse, error)
}
