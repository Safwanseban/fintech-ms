package interfaces

import "fintechGo/internal/types"

type UserInterface interface {
	CreateUser(user *types.AuthUser) error
	FindUserByData(data string) (*types.AuthUser, error)
}
