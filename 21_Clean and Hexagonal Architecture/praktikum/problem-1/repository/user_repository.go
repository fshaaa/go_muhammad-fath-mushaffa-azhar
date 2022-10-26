package repository

import (
	"belajar-go-echo/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user model.User) error
	Find() ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Create(user model.User) error {
	return u.db.Save(&user).Error
}

func (u *userRepository) Find() ([]model.User, error) {
	var user []model.User

	err := u.db.Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
