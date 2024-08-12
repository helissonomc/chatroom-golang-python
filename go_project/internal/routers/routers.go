package routers

import (
	"chatroom/internal/delivery/rest"
	"chatroom/internal/middleware"

	"github.com/gorilla/mux"
)

func InitRouter(userHandler *rest.UserHandler) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users", userHandler.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	return router
}
