package domain

import "time"

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
