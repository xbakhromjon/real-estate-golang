package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"real-estate/domain"
	"strconv"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

func (uc *UserController) Create(context *gin.Context) {
	var request domain.UserCreateRequest
	err := context.Bind(&request)
	if err != nil {
		fmt.Println("Error binding")
		return
	}
	fmt.Println(request)
	response, err := uc.UserUseCase.Create(&request)
	if err != nil {
		fmt.Println("Error processing")
		return
	}
	context.String(http.StatusOK, strconv.FormatInt(response, 10))
}

func (uc *UserController) Update(context *gin.Context) {
	id := context.Params.ByName("id")
	fmt.Println(id)
	var request domain.UserUpdateRequest
	err := context.Bind(&request)
	if err != nil {
		fmt.Println("Error occured binding")
	}
	fmt.Println(request)
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic("converting error")
	}
	err = uc.UserUseCase.Update(ID, &request)
	if err != nil {
		fmt.Println("Error processing")
	}

}

func (uc *UserController) GetOne(context *gin.Context) {
	id := context.Params.ByName("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic("error convert")
	}
	user, err := uc.UserUseCase.GetOne(ID)
	if err != nil {
		panic("error")
	}
	context.JSON(http.StatusOK, user)
}

func (uc *UserController) Block(context *gin.Context) {

}

func (uc *UserController) UnBlock(context *gin.Context) {

}

func (uc *UserController) ChangeStatus(context *gin.Context) {
	id := context.Params.ByName("id")
	status := context.Query("status")
	if len(status) == 0 {
		fmt.Println("query status required")
		return
	}
	ID, _ := strconv.ParseInt(id, 10, 64)
	convertedStatus, _ := strconv.ParseInt(status, 10, 8)
	uc.UserUseCase.ChangeStatus(ID, domain.UserStatus(convertedStatus))

}
