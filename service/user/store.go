package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitub.com/matheus-hrm/curiously/types"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

func ScanRowIntoUser(rows pgx.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password_Hash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserByEmail(email string, c *gin.Context) (*types.User, error) {
	row := s.db.QueryRow(c, "SELECT * FROM users WHERE email = $1", email)
	user := new(types.User)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password_Hash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Fatalf("error scanning row: %s", err)
	}
	return user, nil
}

func (s *Store) GetQuestionsByUserID(id int, c *gin.Context) ([]types.Question, error) {
	rows, err := s.db.Query(c, "SELECT * FROM questions WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	questions := make([]types.Question, 0)
	for rows.Next() {
		question := new(types.Question)
		err := rows.Scan(
			&question.ID,
			&question.UserID,
			&question.Content,
			&question.IsAnonymous,
			&question.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		questions = append(questions, *question)
	}
	return questions, nil
}

func (s *Store) GetUserByID(id int, c *gin.Context) (*types.User, error) {
	row := s.db.QueryRow(c, "SELECT * FROM users WHERE id = $1", id)
	user := new(types.User)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password_Hash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Fatalf("error scanning row: %s", err)
	}

	return user, nil
}

func (s *Store) CreateUser(user types.User, c *gin.Context) error {
	_, err := s.db.Exec(c, "INSERT INTO users ( username,email, password_hash) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password_Hash)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUserByUsername(username string, c *gin.Context) (*types.User, error) {
	row := s.db.QueryRow(c, "SELECT * FROM users WHERE username = $1", username)
	user := new(types.User)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password_Hash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Fatalf("error scanning row: %s", err)
	}
	return user, nil
}
