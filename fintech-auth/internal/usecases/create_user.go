package usecases

import (
	pkg "fintechGo/internal/pkg/middleware"
	"fintechGo/internal/repo/interfaces"
	"fintechGo/internal/types"
	services "fintechGo/internal/usecases/interfaces"
)

type UserUsecase struct {
	userRepo interfaces.AuthRepo
}

// CreateUser implements interfaces.UserInterface
func (r *UserUsecase) CreateUser(user *types.AuthUser) (map[string]string, error) {

	password, err := user.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = password
	if err := r.userRepo.CreateUser(user); err != nil {
		return nil, err
	}
	data, err := pkg.CreateJWT(user.Email)
	if err != nil {
		return nil, err
	}
	return data, nil
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
