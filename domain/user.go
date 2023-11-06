package domain

type UserStatus int64

const (
	UserActive UserStatus = iota
	UserDeleted
	UserBlocked
)

type User struct {
	ID        int64
	Flm       string
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
	ID        int64
	Flm       string
	Phone     string
	CoinCount int32
	Status    int64
}

type UserUseCase interface {
	Create(request UserCreateRequest) (int64, error)
	Update(request UserUpdateRequest) error
	GetOne(ID int64) (UserResponse, error)
	Block(ID int64) error
	UnBlock(ID int64) error
}
