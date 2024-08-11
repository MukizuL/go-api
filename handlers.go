package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Task struct {
	Task_id int `json:"task_id"`
}

func (s *APIServer) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(Task{Task_id: 100})
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handleGetTaskStatus(w http.ResponseWriter, r *http.Request) {

}

func (s *APIServer) handleGetTaskResult(w http.ResponseWriter, r *http.Request) {

}
