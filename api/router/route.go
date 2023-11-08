package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(gin *gin.Engine, db *gorm.DB) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewUserRouter(publicRouter, db)
}
