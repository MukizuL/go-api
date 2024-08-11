package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()

	router.HandleFunc("/task", s.handleCreateTask).Methods("POST") // /task или /task/{task_id}
	router.HandleFunc("/status/{task_id}", s.handleGetTaskStatus).Methods("GET")
	router.HandleFunc("/result/{task_id}", s.handleGetTaskResult).Methods("GET")

	log.Println("Starting server on ", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, router))
}

func main() {
	api := NewAPIServer(":8000", nil)
	api.Serve()
}
