package repository

import (
	"gorm.io/gorm"
	"real-estate/domain"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (u userRepository) Save(user *domain.User) (int64, error) {
	u.database.Save(user)
	return int64(user.ID), nil
}

func (u userRepository) Update(ID int64, user *domain.User) error {
	u.database.Model(user).Updates(user)
	return nil
}

func (u userRepository) GetOne(ID int64) *domain.User {
	var user domain.User
	u.database.First(&user, ID)
	return &user
}

func (u userRepository) ChangeStatus(id int64, status domain.UserStatus) {
	u.database.Model(&domain.User{}).Where("id = ?", id).Update("status", status)
}
