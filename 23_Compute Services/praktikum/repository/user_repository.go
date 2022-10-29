package repository

import (
	"fshaaa-rest-api/model"
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
	var users []model.User

	err := u.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
