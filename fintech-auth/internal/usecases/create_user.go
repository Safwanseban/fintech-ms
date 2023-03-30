package usecases

import (
	"fintechGo/internal/repo/interfaces"
	"fintechGo/internal/types"
	services "fintechGo/internal/usecases/interfaces"
)

type UserUsecase struct {
	userRepo interfaces.AuthRepo
}

// CreateUser implements interfaces.UserInterface
func (r *UserUsecase) CreateUser(user *types.AuthUser) error {
	if err := r.userRepo.CreateUser(user); err != nil {
		return err
	}
	return nil
}

// FindUserByData implements interfaces.UserInterface
func (*UserUsecase) FindUserByData(data string) (*types.AuthUser, error) {
	panic("unimplemented")
}

func NewUserUseCase(user interfaces.AuthRepo) services.UserInterface {
	return &UserUsecase{
		userRepo: user,
	}
}
