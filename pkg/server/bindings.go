package server

import (
	user "github.com/malkhandi-anibrata-tft/self.best-temp/pkg/api/users"
	"gorm.io/gorm"
)

func bindUser(dbHandler *gorm.DB) user.Handler {
	repository := user.NewRepository(dbHandler)
	service := user.NewService(repository)
	return user.NewHandler(service)
}
