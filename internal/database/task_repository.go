package database

import (
	"errors"
	"github.com/talyx/TaskManagerApi/internal/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	if err := r.DB.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) GetTaskById(id uint) (*models.Task, error) {
	var task models.Task
	if err := r.DB.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) UpdateTask(task *models.Task) error {
	if err := r.DB.Save(task).Error; err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) DeleteTaskById(id uint) error {
	if err := r.DB.Delete(&models.Task{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) GetAllTasks() ([]*models.Task, error) {
	var tasks []*models.Task
	if err := r.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetAllTasksByProjectId(projectId uint) ([]*models.Task, error) {
	var tasks []*models.Task
	if err := r.DB.Where("ProjectID = ?", projectId).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
