package types

import (
	"time"

	"github.com/gin-gonic/gin"
)

type UserStorage interface {
	GetUserByEmail(email string, context *gin.Context) (*User, error)
	GetUserByID(id int, context *gin.Context) (*User, error)
	CreateUser(user User, context *gin.Context) error
}

type QuestionStorage interface {
	
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID            int    `json:"id"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	Password_Hash string `json:"-"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt	  time.Time `json:"updated_at"`
}

type Question struct {
	ID          int    `json:"id"`
	Content     string `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	IsAnonymous bool   `json:"is_anonymous"`
}

type Answer struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
