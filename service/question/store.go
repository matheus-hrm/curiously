package question

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitub.com/matheus-hrm/curiously/types"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

func (s *Store) CreateQuestion(payload types.CreateQuestionPayload, c *gin.Context) (*types.Question, error) {
	row := s.db.QueryRow(c,
		"INSERT INTO questions (user_id, content, is_anonymous) VALUES ($1, $2, $3) RETURNING id, user_id, content, is_anonymous::boolean, created_at",
		payload.UserID, payload.Content, payload.IsAnonymous,
	)
	question := new(types.Question)
	err := row.Scan(
		&question.ID,
		&question.UserID,
		&question.Content,
		&question.IsAnonymous,
		&question.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (s *Store) GetQuestionByID(id int, c *gin.Context) (*types.Question, error) {
	row := s.db.QueryRow(c, "SELECT * FROM questions WHERE id = $1", id)
	question := new(types.Question)
	err := row.Scan(
		&question.ID,
		&question.UserID,
		&question.Content,
		&question.IsAnonymous,
		&question.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (s *Store) GetQuestions(c *gin.Context) ([]types.Question, error) {
	rows, err := s.db.Query(c, "SELECT * FROM questions")
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
