package router

import (
	"github.com/gorilla/mux"
	// "../middleware"
)

// Router is exported and used in main.go
// func Router() *mux.Router {
apackage router

import (
	"github.com/abdennour/go-to-do-app/middleware"
	"github.com/gorilla/mux"
)

// Router - is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/healthz", middleware.GetHealth).Methods("GET", "OPTIONS")
	return router
}