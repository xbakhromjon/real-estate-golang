package repository

import (
	"gorm.io/gorm"
	"real-estate/domain"
)

type userRepository struct {
	database gorm.DB
}

func NewUserRepository(db gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (u userRepository) Save(user *domain.User) (int64, error) {
	//TODO implement me
	panic("implement me")
}
