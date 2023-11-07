package usecase

import (
	"fmt"
	"real-estate/domain"
)

type userUseCase struct {
}

func NewUserUseCase() domain.UserUseCase {

	return userUseCase{}
}

func (u userUseCase) Create(request *domain.UserCreateRequest) (int64, error) {
	fmt.Println("2 - marta: ", request)
	return 1, nil
}

func (u userUseCase) Update(request domain.UserUpdateRequest) error {
	return nil
}

func (u userUseCase) GetOne(ID int64) (domain.UserResponse, error) {
	return domain.UserResponse{ID: 1, Flm: "Name", Phone: "998945520609", CoinCount: 12, Status: domain.UserActive}, nil
}

func (u userUseCase) Block(ID int64) error {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) UnBlock(ID int64) error {
	//TODO implement me
	panic("implement me")
}
