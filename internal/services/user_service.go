package services

import (
	"errors"
	"github.com/talyx/TaskManagerApi/internal/database"
	"github.com/talyx/TaskManagerApi/internal/models"
	"github.com/talyx/TaskManagerApi/pkg/logger"
)

type UserService struct {
	UserRepo *database.UserRepository
}

func NewUserService(repo *database.UserRepository) *UserService {
	return &UserService{
		UserRepo: repo,
	}
}

func (s *UserService) CreateUser(name string, email string) (*models.User, error) {
	user := &models.User{
		Names: name,
		Email: email,
	}
	err := s.UserRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserById(id uint) (*models.User, error) {
	user, err := s.UserRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(id uint, name string, email string) (*models.User, error) {
	user, err := s.GetUserById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		logger.Debug("Request to update a non-existent user", map[string]interface{}{"id": id})
		return nil, errors.New("user not found")
	}
	user.Names = name
	user.Email = email
	err = s.UserRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUserById(id uint) error {
	user, err := s.UserRepo.GetUserById(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	return s.UserRepo.DeleteUserById(id)
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	return s.UserRepo.GetAllUsers()
}
