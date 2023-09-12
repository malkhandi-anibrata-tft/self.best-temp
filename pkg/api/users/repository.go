package user

import (
	"fmt"

	goerrors "github.com/go-errors/errors"
	"github.com/malkhandi-anibrata-tft/self.best-temp/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *models.UserDetails) (*models.UserDetails, error)
	GetUserData(id uint64) (user *models.UserDetails, err error)
	SaveUser(user *models.UserDetails) (*models.UserDetails, error)
	DeleteUser(email string) error 
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(user *models.UserDetails) (*models.UserDetails, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, goerrors.New(err)
	}
	fmt.Println("User Added!!")

	return user, nil
}

func (r *repository) GetUserData(id uint64) (user *models.UserDetails, err error) {
	err = r.db.First(&user, "id=?", id).Error
	if err != nil {
		err = goerrors.New(err)
	}

	return
}

// SaveUser implements Repository.
func (r *repository) SaveUser(user *models.UserDetails) (*models.UserDetails, error) {
	err := r.db.Save(user).Error
	if err != nil {
		return nil, goerrors.New(err)
	}

	return user, nil
}

func (r *repository) DeleteUser(email string) error {

	err := r.db.Where("email = ?", email).Delete(&models.UserDetails{}).Error
	if err != nil {
		return goerrors.New(err)
	}

	return nil
}
