package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/talyx/TaskManagerApi/internal/models"
	"github.com/talyx/TaskManagerApi/internal/services"
	"github.com/talyx/TaskManagerApi/pkg/logger"
	"net/http"
	"strconv"
)

type ProjectHandler struct {
	ProjectService *services.ProjectService
	SessionStore   *sessions.FilesystemStore
}

func NewProjectHandler(projecService *services.ProjectService, sessionStore *sessions.FilesystemStore) *ProjectHandler {
	return &ProjectHandler{
		ProjectService: projecService,
		SessionStore:   sessionStore,
	}
}

func (handler *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		logger.Info("create project, bad request", map[string]interface{}{"error": err})
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	logger.Info("create project", map[string]interface{}{"project": project})
	session, err := handler.SessionStore.Get(r, "session")
	if err != nil {
		logger.Error("get session error", map[string]interface{}{"error": err})
		http.Error(w, "BadRequest", http.StatusBadRequest)
		return
	}
	userID := session.Values["UserID"].(uint)
	if err := handler.ProjectService.CreateProject(&project, userID); err != nil {
		logger.Info("create project, server error", map[string]interface{}{"error": err})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("create project successfully", map[string]interface{}{"project": project})
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)

}

func (handler *ProjectHandler) GetProjectById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		logger.Info("get project by id, bad request", map[string]interface{}{"error": err})
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	session, err := handler.SessionStore.Get(r, "session")
	if err != nil {
		logger.Error("get session error", map[string]interface{}{"error": err})
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	userID := session.Values["UserID"].(uint)
	project, err := handler.ProjectService.GetProjectById(userID, uint(id))

	if err != nil {
		logger.Info("get project by id, server error", map[string]interface{}{"error": err})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("get project by id successfully", map[string]interface{}{"project": project})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(project)
}

func (handler *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		logger.Info("update project, bad request", map[string]interface{}{"error": err})
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	session, err := handler.SessionStore.Get(r, "session")
	if err != nil {
		logger.Error("get session error", map[string]interface{}{"error": err})
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	userID := session.Values["UserID"].(uint)
	if err := handler.ProjectService.UpdateProject(&project, userID); err != nil {
		logger.Info("update project, server error", map[string]interface{}{"error": err})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("update project successfully", map[string]interface{}{"project": project})
	w.WriteHeader(http.StatusOK)
}

func (handler *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		logger.Info("delete project, bad request", map[string]interface{}{"error": err})
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	session, err := handler.SessionStore.Get(r, "session")
	if err != nil {
		logger.Error("get session error", map[string]interface{}{"error": err})
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	userID := session.Values["UserID"].(uint)
	if err := handler.ProjectService.DeleteProjectById(userID, uint(id)); err != nil {
		logger.Info("delete project, server error", map[string]interface{}{"error": err})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("delete project successfully", map[string]interface{}{"project": nil})
	w.WriteHeader(http.StatusOK)
}

func (handler *ProjectHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := handler.ProjectService.GetAllProjects()
	if err != nil {
		logger.Info("get all projects, server error", map[string]interface{}{"error": err})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("get all projects successfully", map[string]interface{}{"projects": projects})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}

func (handler *ProjectHandler) GetAllProjectByUserId(w http.ResponseWriter, r *http.Request) {
	session, err := handler.SessionStore.Get(r, "session")
	if err != nil {
		logger.Error("get session error", map[string]interface{}{"error": err})
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	userID := session.Values["UserID"].(uint)
	projects, err := handler.ProjectService.GetAllProjectByUserId(userID)
	if err != nil {
		logger.Info("get all project by user, server error", map[string]interface{}{"error": err})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("get all project by user successfully", map[string]interface{}{"projects": projects})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}
