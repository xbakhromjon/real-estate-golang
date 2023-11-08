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
	err = uc.UserUseCase.Update(request)
	if err != nil {
		fmt.Println("Error processing")
	}

}

func (uc *UserController) GetOne(context *gin.Context) {
	id := context.Params.ByName("id")
	fmt.Println(id)
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
	fmt.Printf("id: %s status: %s", id, status)

}
