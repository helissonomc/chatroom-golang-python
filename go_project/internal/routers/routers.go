package routers

import (
	"chatroom/internal/delivery/rest"
	"chatroom/internal/delivery/wsdelivery"
	"chatroom/internal/middleware"

	"github.com/gorilla/mux"
)

func InitRouter(userHandler *rest.UserHandler, wsHandler *wsdelivery.WebSocketHandler) *mux.Router {
	router := mux.NewRouter()
	// USER
	router.Use(middleware.LoggingMiddleware)
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users", userHandler.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
    
    router.HandleFunc("/ws", wsHandler.HandleWebSocket).Methods("GET")
	return router
}
