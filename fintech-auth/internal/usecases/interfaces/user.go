package interfaces

import "fintechGo/internal/types"

type UserInterface interface {
	CreateUser(user *types.AuthUser) (map[string]string, error)
	FindUserByData(data string) (*types.AuthUser, error)
}
