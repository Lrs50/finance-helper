package repository

import (
	"database/sql"
	"finance-helper/api/internal/models"
	"fmt"

	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(req models.CreateUserRequest) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(
		queryCreateUser,
		req.Username,
		req.Name,
		req.Password,
		req.Email,
		req.PhoneNumber,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Name,
		&user.Password,
		&user.Email,
		&user.PhoneNumber,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil

}

func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(queryGetUserByID, id).Scan(
		&user.ID,
		&user.Username,
		&user.Name,
		&user.Password,
		&user.Email,
		&user.PhoneNumber,
		&user.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	rows, err := r.db.Query(queryGetAllUsers)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Name,
			&user.Password,
			&user.Email,
			&user.PhoneNumber,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}
