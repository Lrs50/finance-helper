package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateUserRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type Transaction struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	Value      float64   `json:"value"`
	CategoryID uuid.UUID `json:"category_id"`
	TagID      uuid.UUID `json:"tag_id"`
	Comment    *string   `json:"comment"` // Pointer because it's nullable
	Date       time.Time `json:"date"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateTransactionRequest struct {
	UserID     uuid.UUID `json:"user_id"`
	Value      float64   `json:"value"`
	CategoryID uuid.UUID `json:"category_id"`
	TagID      uuid.UUID `json:"tag_id"`
	Comment    *string   `json:"comment"`
	Date       time.Time `json:"date"`
}