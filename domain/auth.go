package domain

import "time"

type SignUpRequest struct {
	Flm         string
	PhoneNumber string
	password    string
}

type LoginRequest struct {
	phoneNumber string
	password    string
}

type AccessToken struct {
	ID       int64
	Value    string
	ExpireAt time.Time
}

type AccessTokenResponse struct {
	ID       int64
	Value    string
	ExpireAt time.Time
}

type RefreshToken struct {
	ID       int64
	Value    string
	ExpireAt time.Time
}

type RefreshTokenResponse struct {
	ID       int64
	Value    string
	ExpireAt time.Time
}

type AuthUseCase interface {
	Login(request LoginRequest) (AccessTokenResponse, error)
	Signup(request SignUpRequest) error
}
