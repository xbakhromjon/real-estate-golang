package domain

import "time"

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
