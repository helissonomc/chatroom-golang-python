package repository

// Contains the data access layer, such as database interactions.

import (
	"chatroom/internal/domain"
	"errors"
)

type inMemoryUserRepository struct {
	users map[int]*domain.User
}

func NewInMemoryUserRepository() domain.UserRepository {
	return &inMemoryUserRepository{
		users: make(map[int]*domain.User),
	}
}

func (r *inMemoryUserRepository) Create(user *domain.User) error {
	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.users[user.ID] = user
	return nil
}

func (r *inMemoryUserRepository) GetByID(id int) (*domain.User, error) {
	if user, exists := r.users[id]; exists {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (r *inMemoryUserRepository) Update(user *domain.User) error {
	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}
	r.users[user.ID] = user
	return nil
}

func (r *inMemoryUserRepository) Delete(id int) error {
	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return nil
}

func (r *inMemoryUserRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}
