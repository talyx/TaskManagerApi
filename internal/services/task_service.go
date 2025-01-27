package services

import (
	"errors"
	"github.com/talyx/TaskManagerApi/internal/database"
	"github.com/talyx/TaskManagerApi/internal/models"
	"github.com/talyx/TaskManagerApi/pkg/logger"
)

type TaskService struct {
	TaskRepo    *database.TaskRepository
	ProjectRepo *database.ProjectRepository
}

func NewTaskService(repo *database.TaskRepository, projectRepo *database.ProjectRepository) *TaskService {
	return &TaskService{
		TaskRepo:    repo,
		ProjectRepo: projectRepo,
	}
}

func (t *TaskService) CreateTask(task *models.Task, userID uint) error {
	if task.Title == "" {
		return errors.New("task name can not be empty")
	}
	if task.ProjectID == 0 {
		return errors.New("task must be associated with a project")
	}
	logger.Info("create task: before check", map[string]interface{}{
		"user_id":       userID,
		"taskProjectID": task.ProjectID,
	})
	isAssigned, err := t.ProjectRepo.IsUserAssignedToProject(userID, task.ProjectID)
	if err != nil {
		return err
	}
	if !isAssigned {
		return errors.New("project is not assigned to this user")
	}
	logger.Info("create task: after check", nil)
	return t.TaskRepo.CreateTask(task)
}

func (t *TaskService) GetTaskById(userID, id uint) (*models.Task, error) {
	task, err := t.TaskRepo.GetTaskById(id)
	if err != nil {
		return nil, err
	}
	isAssigned, err := t.ProjectRepo.IsUserAssignedToProject(userID, task.ProjectID)
	if err != nil {
		return nil, err
	}
	if !isAssigned {
		return nil, errors.New("the task refers to a project that belongs to another user")
	}
	return task, nil
}

func (t *TaskService) UpdateTask(task *models.Task, userID uint) error {
	if task.Title == "" {
		return errors.New("task cannot be empty")
	}
	updateTask, err := t.TaskRepo.GetTaskById(task.ID)
	if err != nil {
		return err
	}
	if updateTask == nil {
		return errors.New("task not found")
	}
	if task.ProjectID != updateTask.ProjectID {
		return errors.New("task is not assigned to this project")
	}

	isAssigned, err := t.ProjectRepo.IsUserAssignedToProject(userID, task.ProjectID)
	if err != nil {
		return err
	}
	if !isAssigned {
		return errors.New("the task refers to a project that belongs to another user")
	}

	return t.TaskRepo.UpdateTask(task)
}

func (t *TaskService) DeleteTaskById(userID, id uint) error {
	task, err := t.TaskRepo.GetTaskById(id)
	if err != nil {
		return err
	}
	if task == nil {
		return errors.New("task not found")
	}

	isAssigned, err := t.ProjectRepo.IsUserAssignedToProject(userID, task.ProjectID)
	if err != nil {
		return err
	}
	if !isAssigned {
		return errors.New("the task refers to a project that belongs to another user")
	}
	return t.TaskRepo.DeleteTaskById(id)
}

func (t *TaskService) GetAllTasks() ([]*models.Task, error) {
	return t.TaskRepo.GetAllTasks()
}

func (t *TaskService) GetAllTasksByProjectId(userID, projectID uint) ([]*models.Task, error) {
	isAssigned, err := t.ProjectRepo.IsUserAssignedToProject(userID, projectID)
	if err != nil {
		return nil, err
	}
	if !isAssigned {
		return nil, errors.New("the task refers to a project that belongs to another user")
	}
	return t.TaskRepo.GetAllTasksByProjectId(projectID)
}
