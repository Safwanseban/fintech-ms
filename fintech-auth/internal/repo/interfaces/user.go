package interfaces

import "fintechGo/internal/types"

type AuthRepo interface {
	CreateUser(user *types.AuthUser) error
	GetUserByEmail(email string) (*types.AuthUser, error)
	GetAllUsers() (*[]types.AuthUser, error)
}
