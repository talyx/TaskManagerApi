package services

import (
	"errors"
	"github.com/talyx/TaskManagerApi/internal/database"
	"github.com/talyx/TaskManagerApi/internal/models"
)

type ProjectService struct {
	ProjectRepo *database.ProjectRepository
}

func NewProjectService(projectRepo *database.ProjectRepository) *ProjectService {
	return &ProjectService{ProjectRepo: projectRepo}
}

func (p *ProjectService) CreateProject(project *models.Project) error {
	if project.Name == "" {
		return errors.New("project name can not be empty")
	}
	if project.UserID == 0 {
		return errors.New("project must associated with a user")
	}
	return p.ProjectRepo.CreateProject(project)
}

func (p *ProjectService) GetProjectById(id uint) (*models.Project, error) {
	project, err := p.ProjectRepo.GetProjectById(id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (p *ProjectService) UpdateProject(project *models.Project) error {
	if project.Name == "" {
		return errors.New("project name can not be empty")
	}
	projectUpdated, err := p.ProjectRepo.GetProjectById(project.ID)
	if err != nil {
		return err
	}
	if projectUpdated == nil {
		return errors.New("project doesn't exist")
	}
	return p.ProjectRepo.UpdateProject(project)
}

func (p *ProjectService) DeleteProjectById(id uint) error {
	project, err := p.ProjectRepo.GetProjectById(id)
	if err != nil {
		return err
	}
	if project == nil {
		return errors.New("project doesn't exist")
	}
	return p.ProjectRepo.DeleteProjectById(id)
}

func (p *ProjectService) GetAllProjects() ([]*models.Project, error) {
	return p.ProjectRepo.GetAllProjects()
}

func (p *ProjectService) GetAllProjectByUserId(userId uint) ([]*models.Project, error) {
	return p.ProjectRepo.GetAllProjectByUserId(userId)
}
