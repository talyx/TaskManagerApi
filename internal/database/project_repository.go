package database

import (
	"errors"
	"github.com/talyx/TaskManagerApi/internal/models"
	"github.com/talyx/TaskManagerApi/pkg/logger"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		DB: db,
	}
}

func (r *ProjectRepository) CreateProject(project *models.Project) error {
	if err := r.DB.Create(project).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProjectRepository) GetProjectById(id uint) (*models.Project, error) {
	var project models.Project
	if err := r.DB.First(&project, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &project, nil
}

func (r *ProjectRepository) UpdateProject(project *models.Project) error {
	if err := r.DB.Save(project).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProjectRepository) DeleteProjectById(id uint) error {
	if err := r.DB.Delete(&models.Project{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProjectRepository) GetAllProjects() ([]*models.Project, error) {
	var projects []*models.Project
	if err := r.DB.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepository) GetAllProjectByUserId(id uint) ([]*models.Project, error) {
	var projects []*models.Project
	if err := r.DB.Where("user_id = ?", id).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepository) IsUserAssignedToProject(userID, projectID uint) (bool, error) {
	var count int64
	logger.Info("isAssigned: before request", map[string]interface{}{
		"user_id":    userID,
		"project_id": projectID,
	})
	err := r.DB.
		Table("projects").
		Where("user_id = ? AND ID = ?", userID, projectID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
