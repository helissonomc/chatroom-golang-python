package domain

// Contains the entities and domain-specific logic.
import "fmt"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (user User) String() string {
	return fmt.Sprintf("{ID: %v, Name: %v, Email: %v}", user.ID, user.Name, user.Email)
}
