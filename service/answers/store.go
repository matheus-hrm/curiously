package answers

import (
	"fmt"
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

func (s *Store) CreateAnswer(payload types.CreateAnswerPayload, c *gin.Context) (*types.Answer, error) {
	query := `INSERT INTO answers (question_id, user_id, content) VALUES ($1, $2, $3) RETURNING *`
	return s.executeQueryAndScanAnswer(c, query, payload.QuestionID, payload.UserID, payload.Content)
}

func (s *Store) GetAnswerByID(id int, c *gin.Context) (*types.Answer, error) {
	query := "SELECT * FROM answers WHERE id = $1"
	return s.executeQueryAndScanAnswer(c, query, id)
}

func (s *Store) GetAnswersByQuestionID(id int, c *gin.Context) ([]types.Answer, error) {
	query := "SELECT * FROM answers WHERE question_id = $1"
	return s.executeQueryAndScanAnswers(c, query, id)
}

func (s *Store) executeQueryAndScanAnswer(c *gin.Context, query string, args ...interface{}) (*types.Answer, error) {
	row := s.db.QueryRow(c, query, args...)
	log.Printf("query: %s, args: %v", query, args)
	return scanSingleAnswer(row)
}

func (s *Store) executeQueryAndScanAnswers(c *gin.Context, query string, args ...interface{}) ([]types.Answer, error) {
	rows, err := s.db.Query(c, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	return scanMultipleAnswers(rows)
}

func scanSingleAnswer(row pgx.Row) (*types.Answer, error) {
	answer := new(types.Answer)
	err := row.Scan(
		&answer.ID,
		&answer.QuestionID,
		&answer.UserID,
		&answer.Content,
		&answer.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error scanning row: %w", err)
	}
	return answer, nil
}

func scanMultipleAnswers(rows pgx.Rows) ([]types.Answer, error) {
	answers := make([]types.Answer, 0)
	for rows.Next() {
		answer, err := scanSingleAnswer(rows)
		if err != nil {
			return nil, err
		}
		answers = append(answers, *answer)
	}
	return answers, nil
}
