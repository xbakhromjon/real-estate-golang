package usecase

import (
	"real-estate/domain"
)

type userUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) domain.UserUseCase {

	return userUseCase{userRepository: userRepository}
}

func (u userUseCase) Create(request *domain.UserCreateRequest) (int64, error) {
	id, err := u.userRepository.Save(&domain.User{
		Flm:       request.Flm,
		Phone:     request.Phone,
		CoinCount: 0,
		Status:    0,
	})
	if err != nil {
		panic("error saving")
	}
	return id, nil
}

func (u userUseCase) Update(ID int64, request *domain.UserUpdateRequest) error {
	u.userRepository.Save(&domain.User{
		ID:        ID,
		Flm:       request.Flm,
		Phone:     request.Phone,
		CoinCount: 0,
		Status:    0,
	})
	return nil
}

func (u userUseCase) GetOne(ID int64) (domain.UserResponse, error) {
	user := u.userRepository.GetOne(ID)
	return domain.UserResponse{
		ID:        user.ID,
		Flm:       user.Flm,
		Phone:     user.Phone,
		CoinCount: user.CoinCount,
		Status:    domain.UserStatus(user.Status),
	}, nil
}

func (u userUseCase) Block(ID int64) error {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) UnBlock(ID int64) error {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) ChangeStatus(ID int64, status domain.UserStatus) error {
	u.userRepository.ChangeStatus(ID, status)
	return nil
}
