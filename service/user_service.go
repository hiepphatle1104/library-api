package service

import (
	"context"
	"fmt"
	"library-api/common/utils"
	"library-api/model"
	"library-api/repository"

	"github.com/google/uuid"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateNewUser(user *model.User) error {
	ctx := context.Background()

	// Check exist user
	exists, err := s.repo.ExistByUsername(user.UserName, ctx)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("user exists")
	}

	var newUser *model.User = user

	// Creating userId
	newUser.ID = uuid.New()

	// Hashing password
	hash, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return err
	}

	// Replace password into hash
	newUser.Password = hash

	// Create user
	err = s.repo.CreateUser(newUser, ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) CreateNewSession(user *model.User) error {
	return nil
}

func (s *UserService) ClearCurrSession() error {
	return nil
}
