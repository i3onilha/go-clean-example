package usecase

import (
	"context"
	"time"

	"go-clean-example/internal/dto"
	"go-clean-example/internal/repository"

	"github.com/adityaeka26/go-pkg/logger"
)

type userUsecase struct {
	logger         *logger.Logger
	userRepository repository.UserRepository
}

func NewUserUsecase(logger *logger.Logger, userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		logger:         logger,
		userRepository: userRepo,
	}
}

func (u *userUsecase) GetUserWithOrders(ctx context.Context, userID int32, limit, offset int) (*dto.UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := u.userRepository.GetUserInfo(ctx, userID)
	if err != nil {
		return nil, err
	}
	orders, err := u.userRepository.GetOrdersByUserID(ctx, user.UserID, limit, offset)
	if err != nil {
		return nil, err
	}
	// Orchestration only: delegate to domain entity method for business logic
	ordersResp := make([]dto.Order, len(orders.Orders))
	for i, order := range orders.Orders {
		finalPrice := order.CalculateFinalPrice() // ← business logic in domain layer
		ordersResp[i] = dto.Order{
			UserID:     order.UserID,
			Item:       order.Item,
			Quantity:   order.Quantity,
			Price:      order.Price,
			Discount:   order.DiscountPercent,
			FinalPrice: finalPrice,
		}
	}
	return &dto.UserResponse{
		UserID:   user.UserID,
		Name:     user.Name,
		Email:    user.Email,
		Location: user.Location,
		Orders:   ordersResp,
	}, nil
}
