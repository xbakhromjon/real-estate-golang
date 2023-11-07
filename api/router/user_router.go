package router

import (
	"github.com/gin-gonic/gin"
	"real-estate/api/controller"
	"real-estate/usecase"
)

func NewUserRouter(group *gin.RouterGroup) {
	uc := controller.UserController{UserUseCase: usecase.NewUserUseCase()}
	group.GET("/users/:id", uc.GetOne)
	group.POST("/users", uc.Create)
	group.PUT("/users/:id", uc.Update)
	group.PATCH("/users/:id/status", uc.ChangeStatus)
}
