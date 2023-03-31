package interfaces

import "fintechGo/internal/types"

type UserInterface interface {
	CreateUser(user *types.AuthUser) error
	ValidateUser(user *types.AuthUser) (map[string]string, error)
	//FindAllUsers() (*[]types.AuthUser, error)
}
