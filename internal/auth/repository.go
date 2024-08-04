package auth

import (
	"database/sql"
	"errors"
)

type Repository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
}

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{DB: db}
}

func (r *PostgresRepository) CreateUser(user *User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	query := `SELECT id, username, email, password FROM users WHERE email = $1`
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
