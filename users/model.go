package users

import (
	"time"
	"uuid"
)

type User struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Fullname   string    `json:"full_name" gorm:"not null"`
	Email      string    `json:"email" gorm:"not null, uniqueIndex"`
	Password   string    `json:"password" gorm:"not null"`
	Created_at time.Time `json:"created_at" gorm:"not null"`
	Updated_at time.Time `json:"update_at" gorm:"not null"`
}

// CreateUserRequest adalah body JSON untuk POST /user
type CreateUserRequest struct {
	Fullname   string    `json:"full_name" binding:"required,min=3"`
	Email      string    `json:"email" binding:"required,email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"update_at"`
}

// UpdateTodoRequest adalah body JSON untuk PUT /todos/:id
type UpdateUserRequest struct {
	Fullname   string    `json:"full_name" binding:"required,min=3"`
	Email      string    `json:"email" binding:"required,email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"update_at"`
}
