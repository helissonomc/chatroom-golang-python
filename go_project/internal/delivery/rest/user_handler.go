package rest

// Contains the delivery mechanism (e.g., HTTP handlers, gRPC services).

import (
	"chatroom/internal/domain"
	"chatroom/internal/usecase"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: uc,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("aaaaa", user)
	if err := h.UserUsecase.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("aaaa", params)
	id, err := strconv.Atoi(params["id"])
	fmt.Println("bbbb", id)

	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.UserUsecase.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, _ := h.UserUsecase.GetAll()
	json.NewEncoder(w).Encode(users)
}
