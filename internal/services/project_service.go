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
	return p.ProjectRepo.CreateProject(project)
}

func (p *ProjectService) GetProjectById(id uint) (*models.Project, error) {
	project, err := p.ProjectRepo.GetProjectById(id)
	if err != nil {
		return nil, err
	}
	if project == nil {
		return nil, errors.New("project not found")
	}
	return project, nil
}

func (p *ProjectService) UpdateProject(project *models.Project) error {
	if project.Name == "" {
		return errors.New("project name can not be empty")
	}
	return p.ProjectRepo.UpdateProject(project)
}

func (p *ProjectService) DeleteProjectById(id uint) error {
	return p.ProjectRepo.DeleteProjectById(id)
}

func (p *ProjectService) GetAllProjects() ([]models.Project, error) {
	return p.ProjectRepo.GetAllProjects()
}
