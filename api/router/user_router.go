package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"real-estate/api/controller"
	"real-estate/repository"
	"real-estate/usecase"
)

func NewUserRouter(group *gin.RouterGroup, db gorm.DB) {
	uc := controller.UserController{UserUseCase: usecase.NewUserUseCase(repository.NewUserRepository(db))}
	group.GET("/users/:id", uc.GetOne)
	group.POST("/users", uc.Create)
	group.PUT("/users/:id", uc.Update)
	group.PATCH("/users/:id/status", uc.ChangeStatus)
}
