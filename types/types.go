package types

import (
	"time"

	"github.com/gin-gonic/gin"
)

type UserStorage interface {
	GetUserByEmail(email string, context *gin.Context) (*User, error)
	GetUserByID(id int, context *gin.Context) (*User, error)
	CreateUser(user User, context *gin.Context) error
	GetQuestionsByUserID(id int, context *gin.Context) ([]Question, error)
	GetUserByUsername(username string, context *gin.Context) (*User, error)
}

type QuestionStorage interface {
	CreateQuestion(payload CreateQuestionPayload, context *gin.Context) (*Question, error)
	GetQuestionByID(id int, context *gin.Context) (*Question, error)
	GetQuestions(context *gin.Context) ([]Question, error)
}

type AnswerStorage interface {
	CreateAnswer(payload CreateAnswerPayload, context *gin.Context) (*Answer, error)
	GetAnswerByID(id int, context *gin.Context) (*Answer, error)
	GetAnswersByQuestionID(id int, context *gin.Context) ([]Answer, error)
}

type CreateAnswerPayload struct {
	Content    string    `json:"content" validate:"required"`
	UserID     int       `json:"userid"`
	QuestionID int       `json:"questionid"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateQuestionPayload struct {
	Content     string    `json:"content" validate:"required"`
	IsAnonymous bool      `json:"is_anonymous"`
	UserID      int       `json:"userid"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserProfile struct {
	Username  string            `json:"username"`
	Email     string            `json:"email"`
	Questions []ProfileQuestion `json:"questions"`
	CreatedAt time.Time         `json:"created_at"`
}

type ProfileQuestion struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Answered  bool      `json:"answered"`
	Answer    []string  `json:"answer"`
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
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	Password_Hash string    `json:"-"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Question struct {
	ID          int       `json:"id"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      int       `json:"user_id"`
	IsAnonymous bool      `json:"is_anonymous"`
}

type Answer struct {
	ID         int       `json:"id"`
	QuestionID int       `json:"question_id"`
	UserID     int       `json:"user_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}
