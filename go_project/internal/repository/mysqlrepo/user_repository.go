package mysqlrepo

import (
	"chatroom/internal/domain"
	"database/sql"
	"errors"
)

type mysqlUserRepository struct {
	dbClient *DbClient
}

func NewMySQLUserRepository(dbClient *DbClient) domain.UserRepository {
	return &mysqlUserRepository{dbClient: dbClient}
}

func (r *mysqlUserRepository) Create(user *domain.User) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := r.dbClient.DB.Exec(query, user.Name, user.Email)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)
	return nil
}

func (r *mysqlUserRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User
	query := "SELECT id, name, email FROM users"
	rows, err := r.dbClient.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *mysqlUserRepository) GetByID(id int) (*domain.User, error) {
	user := &domain.User{}
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.dbClient.DB.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (r *mysqlUserRepository) Update(user *domain.User) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := r.dbClient.DB.Exec(query, user.Name, user.Email, user.ID)
	return err
}

func (r *mysqlUserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.dbClient.DB.Exec(query, id)
	return err
}
