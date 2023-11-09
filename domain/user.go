package domain

import (
	"gorm.io/gorm"
	"time"
)

type UserStatus int8

const (
	UserActive UserStatus = iota
	UserDeleted
	UserBlocked
)

type User struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Flm       string         `gorm:"not null"`
	Phone     string
	CoinCount int32
	Status    int64
}

type UserCreateRequest struct {
	Flm   string
	Phone string
}

type UserUpdateRequest struct {
	Flm   string
	Phone string
}

type UserResponse struct {
	ID        int64      `json:"id"`
	Flm       string     `json:"flm"`
	Phone     string     `json:"phone"`
	CoinCount int32      `json:"coinCount"`
	Status    UserStatus `json:"status"`
}

type UserUseCase interface {
	Create(request *UserCreateRequest) (int64, error)
	Update(ID int64, request *UserUpdateRequest) error
	GetOne(ID int64) (UserResponse, error)
	Block(ID int64) error
	UnBlock(ID int64) error
	ChangeStatus(ID int64, status UserStatus) error
}

type UserRepository interface {
	Save(user *User) (int64, error)
	Update(ID int64, user *User) error
	GetOne(ID int64) *User
	ChangeStatus(id int64, status UserStatus)
}
