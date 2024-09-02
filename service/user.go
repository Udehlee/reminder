package service

import (
	"fmt"

	"github.com/Udehlee/payment-reminder/db/db"
	"github.com/Udehlee/payment-reminder/internals"
	models "github.com/Udehlee/payment-reminder/models/user"
	"github.com/Udehlee/payment-reminder/utils"
)

type Service struct {
	store     db.Store
	logger    internals.Logger
	scheduler *internals.Scheduler
}

func NewService(db db.Store, logger internals.Logger, scheduler *internals.Scheduler) *Service {
	return &Service{
		store:     db,
		logger:    logger,
		scheduler: scheduler,
	}
}

// CreateUser creates a new user
func (s *Service) CreateUser(fname, lname, email, password string) (models.User, error) {
	s.logger.Info("Starting user creation process")

	hashedPwd, err := utils.HashPassword(password)
	if err != nil {
		s.logger.Error("Error hashing password")
		return models.User{}, fmt.Errorf("error hashing password %s", err)
	}

	user := models.User{
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Password:  hashedPwd,
	}

	if err := s.store.SaveUser(user); err != nil {
		s.logger.Error("couldnt save user")
		return models.User{}, fmt.Errorf("error saving to the database %s", err)
	}

	return user, nil
}

// CheckUser confirms user details
func (s Service) CheckUser(email, password string) (models.User, error) {
	s.logger.Info("checking user email")

	user, err := s.store.UserEmail(email)
	if err != nil {
		s.logger.Error("user email not found")
		return models.User{}, fmt.Errorf("user not found")
	}

	err = utils.ComparePasswordHash(user.Password, password)
	if err != nil {
		return models.User{}, fmt.Errorf("wrong password")
	}

	return user, nil
}
