package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/talyx/TaskManagerApi/internal/database"
	"github.com/talyx/TaskManagerApi/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	DB         *gorm.DB
	Repository *database.UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	// Инициализация in-memory SQLite для тестов
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.NoError(err)

	// Миграция таблиц
	err = db.AutoMigrate(&models.User{})
	suite.NoError(err)

	suite.DB = db
	suite.Repository = &database.UserRepository{DB: db}
}

func (suite *UserRepositoryTestSuite) TestCreateUser() {
	user := &models.User{Names: "John Doe", Email: "john.doe@example.com"}
	err := suite.Repository.CreateUser(user)
	suite.NoError(err)
	assert.NotZero(suite.T(), user.ID)
}

func (suite *UserRepositoryTestSuite) TestGetUserByID() {
	user := &models.User{Names: "Jane Doe", Email: "jane.doe@example.com"}
	err := suite.Repository.CreateUser(user)
	suite.NoError(err)

	fetchedUser, err := suite.Repository.GetUserById(user.ID)
	suite.NoError(err)
	assert.Equal(suite.T(), user.Email, fetchedUser.Email)
}

func (suite *UserRepositoryTestSuite) TestUpdateUser() {
	user := &models.User{Names: "Jane Doe", Email: "jane.doe@example.com"}
	err := suite.Repository.CreateUser(user)
	suite.NoError(err)

	user.Names = "Jane Smith"
	err = suite.Repository.UpdateUser(user)
	suite.NoError(err)

	updatedUser, err := suite.Repository.GetUserById(user.ID)
	suite.NoError(err)
	assert.Equal(suite.T(), "Jane Smith", updatedUser.Names)
}

func (suite *UserRepositoryTestSuite) TestDeleteUser() {
	user := &models.User{Names: "John Doe", Email: "john.doe@example.com"}
	err := suite.Repository.CreateUser(user)
	suite.NoError(err)

	err = suite.Repository.DeleteUserById(user.ID)
	suite.NoError(err)

	deletedUser, err := suite.Repository.GetUserById(user.ID)
	suite.NoError(err)
	assert.Nil(suite.T(), deletedUser)
}

func (suite *UserRepositoryTestSuite) TestGetAllUsers() {
	users := []models.User{
		{Names: "User1", Email: "user1@example.com"},
		{Names: "User2", Email: "user2@example.com"},
	}

	for _, user := range users {
		err := suite.Repository.CreateUser(&user)
		suite.NoError(err)
	}

	allUsers, err := suite.Repository.GetAllUsers()
	suite.NoError(err)
	assert.Len(suite.T(), allUsers, 2)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
