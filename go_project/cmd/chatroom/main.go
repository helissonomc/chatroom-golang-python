package main

import (
	"chatroom/internal/delivery/rest"
	"chatroom/internal/delivery/wsdelivery"
	"chatroom/internal/repository/mysqlrepo"
	"chatroom/internal/repository/wsrepo"
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

    // Initialize the repository, use case, and handler layers
    repo := wsrepo.NewInMemoryConnectionRepo()
    wsUsecase := usecase.NewWebSocketUsecase(repo)
    wsHandler := wsdelivery.NewWebSocketHandler(wsUsecase)


	router := routers.InitRouter(userHandler, wsHandler)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
