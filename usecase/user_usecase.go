package usecase

import "real-estate/domain"

type userUseCase struct {
}

func NewUserUseCase() domain.UserUseCase {

	return userUseCase{}
}

func (u userUseCase) Create(request domain.UserCreateRequest) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) Update(request domain.UserUpdateRequest) error {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) GetOne(ID int64) (domain.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) Block(ID int64) error {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) UnBlock(ID int64) error {
	//TODO implement me
	panic("implement me")
}
