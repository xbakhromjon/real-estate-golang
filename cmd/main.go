package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"real-estate/api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	dsn := "host=localhost user=postgres password=123 dbname=real_estate_go port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect DB")
	}
	router.Setup(gin, db)
	gin.Run(":8080")
}
