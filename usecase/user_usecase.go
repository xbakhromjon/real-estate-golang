package usecase

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"real-estate/domain"
)

type userUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) domain.UserUseCase {

	return userUseCase{userRepository: userRepository}
}

func (u userUseCase) Create(request *domain.UserCreateRequest) (int64, error) {
	dsn := "host=localhost user=postgres password=123 dbname=real_estate_go port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect DB")
	}
	// Migrate the schema
	db.AutoMigrate(&domain.User{})

	// create
	db.Create(&domain.User{
		Flm:       "FLM",
		Phone:     "998945520609",
		CoinCount: 10,
		Status:    0,
	})

	var user domain.User
	db.First(&user, 2)
	fmt.Println(user)

	db.Model(&user).Updates(domain.User{Flm: "Test", CoinCount: 100})

	db.Delete(&user)
	u.userRepository.Save(&user)
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

func (u userUseCase) ChangeStatus(ID int64, status domain.UserStatus) error {
	//TODO implement me
	panic("implement me")
}
