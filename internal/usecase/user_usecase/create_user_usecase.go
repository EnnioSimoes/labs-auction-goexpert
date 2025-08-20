package user_usecase

import (
	"context"
	"fullcycle-auction_go/internal/entity/user_entity"
	"fullcycle-auction_go/internal/internal_error"

	"github.com/google/uuid"
)

func (u *UserUseCase) CreateUser(
	ctx context.Context, user UserInputDTO) *internal_error.InternalError {
	userEntity := &user_entity.User{
		Id:   uuid.New().String(),
		Name: user.Name,
	}

	if err := u.UserRepository.CreateUser(ctx, userEntity); err != nil {
		return err
	}

	return nil
}
