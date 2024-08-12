package domain

// Contains the entities and domain-specific logic.
type UserRepository interface {
	Create(user *User) error
	GetByID(id int) (*User, error)
	GetAll() ([]*User, error)
	Update(user *User) error
	Delete(id int) error
}
