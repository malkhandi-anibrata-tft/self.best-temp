package server

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Setup(dbHandler *gorm.DB) {
	fmt.Println("Welcome to the server")

	e := echo.New()

	//handlers
	userHandler := bindUser(dbHandler)
	e.Validator = &CustomValidator{validator: validator.New()}
	user := e.Group("/user")
	user.GET("/info", userHandler.GetUserInfo)
	user.POST("/create", userHandler.CreateUserInfo)
	user.PUT("/update", userHandler.UpdateDetails)
	user.DELETE("/delete", userHandler.DeleteUser)

	e.Start(viper.GetString("server.port"))
}
