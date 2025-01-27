package services

import (
	"errors"
	"github.com/talyx/TaskManagerApi/internal/database"
	"github.com/talyx/TaskManagerApi/internal/models"
	"github.com/talyx/TaskManagerApi/internal/utils"
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

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	var err error
	logger.Debug("Password befor hashing", map[string]interface{}{
		"user_password": user.PasswordHash,
	})
	user.PasswordHash, err = utils.HashPassword(user.PasswordHash)
	logger.Debug("Generated hash", map[string]interface{}{"hash": user.PasswordHash})
	if err != nil {
		return nil, err
	}
	err = s.UserRepo.CreateUser(user)
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

func (s *UserService) GetUserByLogin(login string) (*models.User, error) {
	return s.UserRepo.GetUserByLogin(login)
}
