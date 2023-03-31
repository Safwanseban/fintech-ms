package usecases

import (
	"errors"
	pkg "fintechGo/internal/pkg/middleware"
	"fintechGo/internal/repo/interfaces"
	"fintechGo/internal/types"
	services "fintechGo/internal/usecases/interfaces"

	"github.com/rs/zerolog"
)

type UserUsecase struct {
	userRepo interfaces.AuthRepo
	logger   *zerolog.Logger
}

// CreateUser implements interfaces.UserInterface
func (r *UserUsecase) CreateUser(user *types.AuthUser) error {

	password, err := user.HashPassword(user.Password)
	if err != nil {
		r.logger.Error().Err(err).Send()
		return err
	}

	user.Password = password
	if err := r.userRepo.CreateUser(user); err != nil {
		r.logger.Error().Send()
		return err
	}
	return nil
}

// FindUserByData implements interfaces.UserInterface
func (r *UserUsecase) ValidateUser(user *types.AuthUser) (map[string]string, error) {
	currentPassword := user.Password
	user, err := r.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		r.logger.Error().Err(err).Send()
		return nil, err
	}
	if ok := user.CheckPassword(user.Password, currentPassword); !ok {
		r.logger.Error().Err(errors.New("password doesn't match")).Send()
		return nil, errors.New("password doesn't match")
	}
	jwt, err := pkg.CreateJWT(user.Email)
	if err != nil {
		r.logger.Error().Send()
		return nil, err
	}
	return jwt, nil

}

func NewUserUseCase(user interfaces.AuthRepo, logger *zerolog.Logger) services.UserInterface {
	return &UserUsecase{
		userRepo: user,
		logger:   logger,
	}
}
