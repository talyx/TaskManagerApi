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

type TaskHandler struct {
	TaskService  *services.TaskService
	SessionStore *sessions.FilesystemStore
}

func NewTaskHandler(taskService *services.TaskService, sessionStore *sessions.FilesystemStore) *TaskHandler {
	return &TaskHandler{TaskService: taskService,
		SessionStore: sessionStore}
}

func (th *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		logger.Error("invalid input, request body error", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "Invalid input, request body error", http.StatusBadRequest)
		return
	}
	session, err := th.SessionStore.Get(r, "session")
	if err != nil {
		logger.Error("get session error", map[string]interface{}{"error": err})
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	userID := session.Values["UserID"].(uint)
	logger.Debug("create task", map[string]interface{}{
		"task": task,
	})
	if err := th.TaskService.CreateTask(&task, userID); err != nil {
		logger.Error("cannot create task, server error", map[string]interface{}{
			"error": err,
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("create task", map[string]interface{}{
		"task": task,
	})
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (th *TaskHandler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		logger.Error("invalid input, request id error", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "invalid input, request id error", http.StatusBadRequest)
		return
	}
	session, err := th.SessionStore.Get(r, "session")
	if err != nil {
		logger.Error("get session error", map[string]interface{}{"error": err})
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	userID := session.Values["UserID"].(uint)
	task, err := th.TaskService.GetTaskById(userID, uint(id))
	if err != nil {
		logger.Error("cannot get task, server error", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	if task == nil {
		logger.Error("cannot get task, server error", nil)
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	logger.Info("get task", map[string]interface{}{
		"task": task,
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (th *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		logger.Error("invalid input, request body error", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "invalid input, request body error", http.StatusBadRequest)
		return
	}
	session, err := th.SessionStore.Get(r, "session")
	if err != nil {
		logger.Error("get session error", map[string]interface{}{"error": err})
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	userID := session.Values["UserID"].(uint)
	if err := th.TaskService.UpdateTask(&task, userID); err != nil {
		logger.Error("cannot update task, server error", map[string]interface{}{
			"error": err,
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("task update successfully", map[string]interface{}{
		"task": task,
	})
	w.WriteHeader(http.StatusOK)
}

func (th *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		logger.Error("bed request", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	session, err := th.SessionStore.Get(r, "session")
	if err != nil {
		logger.Error("get session error", map[string]interface{}{"error": err})
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	userID := session.Values["UserID"].(uint)

	err = th.TaskService.DeleteTaskById(userID, uint(id))
	if err != nil {
		logger.Error("delete task error", map[string]interface{}{
			"error": err,
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("delete task successfully", map[string]interface{}{
		"task": id,
	})
	w.WriteHeader(http.StatusOK)
}

func (th *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.TaskService.GetAllTasks()
	if err != nil {
		logger.Error("cannot get all tasks", map[string]interface{}{
			"error": err,
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("get all tasks successfully", map[string]interface{}{
		"tasks": tasks,
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (th *TaskHandler) GetAllTasksByProjectId(w http.ResponseWriter, r *http.Request) {
	projectId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		logger.Error("invalid input, request id error", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "invalid input, request id error", http.StatusBadRequest)
		return
	}
	session, err := th.SessionStore.Get(r, "session")
	if err != nil {
		logger.Error("get session error", map[string]interface{}{"error": err})
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	userID := session.Values["UserID"].(uint)
	projects, err := th.TaskService.GetAllTasksByProjectId(userID, uint(projectId))
	if err != nil {
		logger.Error("get all task by project id error, server error", map[string]interface{}{"id": projectId})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("success get all tasks by project id", nil)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}
