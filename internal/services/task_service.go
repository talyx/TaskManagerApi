package services

import (
	"errors"
	"github.com/talyx/TaskManagerApi/internal/database"
	"github.com/talyx/TaskManagerApi/internal/models"
)

type TaskService struct {
	TaskRepo *database.TaskRepository
}

func NewTaskService(repo *database.TaskRepository) *TaskService {
	return &TaskService{
		TaskRepo: repo,
	}
}

func (t *TaskService) CreateTask(task *models.Task) error {
	if task.Title == "" {
		return errors.New("task name can not be empty")
	}
	if task.ProjectID == 0 {
		return errors.New("task must be associated with a project")
	}
	return t.TaskRepo.CreateTask(task)
}

func (t *TaskService) GetTaskById(id uint) (*models.Task, error) {
	task, err := t.TaskRepo.GetTaskById(id)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (t *TaskService) UpdateTask(task *models.Task) error {
	if task.Title == "" {
		return errors.New("task cannot be empty")
	}
	return t.UpdateTask(task)
}

func (t *TaskService) DeleteTaskById(id uint) error {
	return t.DeleteTaskById(id)
}

func (t *TaskService) GetAllTasks() ([]models.Task, error) {
	return t.TaskRepo.GetAllTasks()
}
