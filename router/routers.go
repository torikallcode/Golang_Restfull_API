package router

import (
	"golang_restfull_API/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	routers := mux.NewRouter()

	routers.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	routers.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	routers.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	routers.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	routers.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	return routers
}
