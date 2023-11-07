package usecase

import "real-estate/domain"

type authUseCase struct {
}

func NewAuthUseCase() domain.AuthUseCase {

	return &authUseCase{}
}

func (a authUseCase) Login(request domain.LoginRequest) (domain.AccessTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a authUseCase) Signup(request domain.SignUpRequest) error {
	//TODO implement me
	panic("implement me")
}
