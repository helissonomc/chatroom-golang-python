package main

import (
	"chatroom/internal/delivery/rest"
	"chatroom/internal/repository/mysqlrepo"
	"chatroom/internal/routers"
	"chatroom/internal/usecase"
	"log"
	"net/http"
)

func main() {
	dbClient := mysqlrepo.InitDB()
	defer dbClient.DB.Close()

	userRepo := mysqlrepo.NewMySQLUserRepository(dbClient)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := rest.NewUserHandler(userUsecase)
	router := routers.InitRouter(userHandler)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
