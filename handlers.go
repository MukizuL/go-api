package main

import (
	"encoding/json"
	"net/http"
)

type TaskService struct {
	store Storage
}

type Task struct {
	Task_id string `json:"task_id"`
	Status  string `json:"status"`
	Result  string `json:"result"`
}

func NewTaskService(s Storage) *TaskService {
	return &TaskService{store: s}
}

func (s *TaskService) RegisterRoutes() {
	http.HandleFunc("POST /task", s.handleCreateTask)
	http.HandleFunc("GET /status/{task_id}", s.handleGetTaskStatus)
	http.HandleFunc("GET /result/{task_id}", s.handleGetTaskResult)
}

// WriteJSON Takes in ResponseWriter, status code and any data and encode it as JSON
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// handleCreateTask takes Task in JSON, adds it to database and returns its id
func (s *TaskService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	return
	//}
	//defer r.Body.Close()

	var task Task
	//err = json.Unmarshal(body, &task)
	//if err != nil {
	//	WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Wrong payload."})
	//	return
	//}

	id, err := s.store.CreateTask(task)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Failed to create task."})
		return
	}
	WriteJSON(w, http.StatusCreated, Task{Task_id: id})
}

// handleGetTaskStatus takes Task id and searches for it in db. If it finds, it returns its status
func (s *TaskService) handleGetTaskStatus(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("task_id")
	if id == "" {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Task ID cannot be empty."})
		return
	}

	task, err := s.store.GetTask(id)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, errNotFound)
		return
	}
	WriteJSON(w, http.StatusOK, Task{Status: task.Status})
}

// handleGetTaskResult takes Task id and searches for it in db. If it finds, it returns its result
func (s *TaskService) handleGetTaskResult(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("task_id")
	if id == "" {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Task ID cannot be empty."})
		return
	}

	task, err := s.store.GetTask(id)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, errNotFound)
		return
	}
	WriteJSON(w, http.StatusOK, Task{Status: task.Result})
}
