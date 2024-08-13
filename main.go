package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewRaiStorage()
	taskService := NewTaskService(store)
	taskService.RegisterRoutes()

	log.Println("Starting server on ", ":8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
