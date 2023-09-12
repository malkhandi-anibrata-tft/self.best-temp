package user

import (
	"time"

	"github.com/malkhandi-anibrata-tft/self.best-temp/models"
)

type Service interface {
	CreateUser(req *CreateDetailsRequest) (*models.UserDetails, error)
	GetUserInfo(userid uint64) (UserDetailsinfo, error)
	UpdateDetails(req *UpdateDetailsRequest) (*models.UserDetails, error)
	DeleteUser(req *DeleteUserRequest) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateUser(req *CreateDetailsRequest) (*models.UserDetails, error) {
	user := new(models.UserDetails)
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.Password = req.Password
	user.CreatedAt = time.Now()

	user, err := s.repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetUserInfo(userid uint64) (UserDetailsinfo, error) {
	getuserdata, err := s.repository.GetUserData(userid)
	if err != nil {
		return UserDetailsinfo{}, err
	}

	user := UserDetailsinfo{
		UserId:    getuserdata.Id,
		FirstName: getuserdata.FirstName,
		LastName:  getuserdata.LastName,
		Email:     getuserdata.Email,
		Password:  getuserdata.Password,
	}

	return user, nil
}

func (s *service) UpdateDetails(req *UpdateDetailsRequest) (*models.UserDetails, error) {
	var user *models.UserDetails
	var err error

	user, err = s.repository.GetUserData(req.UserId)
	if err != nil {
		return nil, err
	}

	if len(req.FirstName) != 0 {
		user.FirstName = req.FirstName
	}

	if len(req.LastName) != 0 {
		user.LastName = req.LastName
	}

	if len(req.Email) != 0 {
		user.Email = req.Email
	}

	if len(req.Password) != 0 {
		user.Password = req.Password
	}

	user, err = s.repository.SaveUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *service) DeleteUser(req *DeleteUserRequest) error {

	err := s.repository.DeleteUser(req.Email)
	if err != nil {
		return err
	}
	return nil
}
