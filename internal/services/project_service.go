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

func (p *ProjectService) CreateProject(project *models.Project, userID uint) error {
	if project.Name == "" {
		return errors.New("project name can not be empty")
	}
	if userID == 0 {
		return errors.New("project must associated with a user")
	}
	project.UserID = userID
	return p.ProjectRepo.CreateProject(project)
}

func (p *ProjectService) GetProjectById(userID, id uint) (*models.Project, error) {
	project, err := p.ProjectRepo.GetProjectById(id)
	if err != nil {
		return nil, err
	}
	isAssigned, err := p.ProjectRepo.IsUserAssignedToProject(userID, id)
	if err != nil {
		return nil, err
	}
	if !isAssigned {
		return nil, errors.New("project is not assigned to this user")
	}
	return project, nil
}

func (p *ProjectService) UpdateProject(project *models.Project, userID uint) error {
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
	isAssigned, err := p.ProjectRepo.IsUserAssignedToProject(userID, project.ID)
	if err != nil {
		return err
	}
	if !isAssigned {
		return errors.New("project is not assigned to this user")
	}
	return p.ProjectRepo.UpdateProject(project)
}

func (p *ProjectService) DeleteProjectById(userID, id uint) error {
	project, err := p.ProjectRepo.GetProjectById(id)
	if err != nil {
		return err
	}
	if project == nil {
		return errors.New("project doesn't exist")
	}
	isAssigned, err := p.ProjectRepo.IsUserAssignedToProject(userID, project.ID)
	if err != nil {
		return err
	}
	if !isAssigned {
		return errors.New("project is not assigned to this user")
	}
	return p.ProjectRepo.DeleteProjectById(id)
}

func (p *ProjectService) GetAllProjects() ([]*models.Project, error) {
	return p.ProjectRepo.GetAllProjects()
}

func (p *ProjectService) GetAllProjectByUserId(userId uint) ([]*models.Project, error) {
	return p.ProjectRepo.GetAllProjectByUserId(userId)
}

func (p *ProjectService) IsUserAssignedToProject(projectId uint, userId uint) (bool, error) {
	return p.ProjectRepo.IsUserAssignedToProject(projectId, userId)
}
