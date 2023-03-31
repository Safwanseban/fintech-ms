package repo

import (
	"fintechGo/internal/repo/interfaces"
	"fintechGo/internal/types"

	"gorm.io/gorm"
)

type Db struct {
	db *gorm.DB
}

// GetAllUsers implements interfaces.AuthRepo
// func (Db *Db) GetAllUsers() (*[]types.AuthUser, error) {
// 	var users *[]types.AuthUser

// 	result := Db.db.Find(&users)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return users, nil

// }

// GetUserByEmail implements interfaces.AuthRepo
func (Db *Db) GetUserByEmail(email string) (*types.AuthUser, error) {
	var user types.AuthUser
	result := Db.db.Where("email=?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

// CreateUser implements interfaces.AuthRepo
func (Db *Db) CreateUser(user *types.AuthUser) error {

	result := Db.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return gorm.ErrInvalidField
	}
	return nil
}

func NewUser(DB *gorm.DB) interfaces.AuthRepo {

	return &Db{
		db: DB,
	}
}
