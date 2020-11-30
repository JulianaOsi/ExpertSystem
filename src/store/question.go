package store

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v4"
)

type Question struct {
	Id        int
	Text      string
	SymptomId int
}

func (s *Store) CreateOrUpdateQuestion(ctx context.Context, question *Question) error {
	sql, _, err := goqu.Insert("question").
		Rows(goqu.Record{
			"id":         question.Id,
			"text":       question.Text,
			"symptom_id": question.SymptomId,
		}).
		OnConflict(goqu.DoUpdate("id", goqu.Record{
			"text":       question.Text,
			"symptom_id": question.SymptomId,
		})).ToSQL()
	if err != nil {
		return fmt.Errorf("sql query build failed: %v", err)
	}
	if _, err := s.connPool.Exec(ctx, sql); err != nil {
		return fmt.Errorf("execute a query failed: %v", err)
	}

	return nil
}

func (s *Store) GetAllQuestions(ctx context.Context) ([]*Question, error) {
	sql, _, err := goqu.Select().From("question").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("sql query build failed: %v", err)
	}

	rows, err := s.connPool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("execute a query failed: %v", err)
	}
	defer rows.Close()

	var questions []*Question

	for rows.Next() {
		question, err := readQuestion(rows)
		if err != nil {
			return nil, fmt.Errorf("read question failed: %v", question)
		}
		questions = append(questions, question)
	}

	return questions, nil
}

func (s *Store) GetQuestionsBySymptomId(ctx context.Context, symptomId int) ([]*Question, error) {
	sql, _, err := goqu.Select().
		From("question").
		Where(goqu.C("symptom_id").Eq(symptomId)).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("sql query build failed: %v", err)
	}

	rows, err := s.connPool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("execute a query failed: %v", err)
	}
	defer rows.Close()

	var questions []*Question

	for rows.Next() {
		question, err := readQuestion(rows)
		if err != nil {
			return nil, fmt.Errorf("read question failed: %v", question)
		}
		questions = append(questions, question)
	}

	return questions, nil
}

func (s *Store) GetQuestionsById(ctx context.Context, id *int) (*Question, error) {
	sql, _, err := goqu.Select().
		From("question").
		Where(goqu.C("id").Eq(id)).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("sql query build failed: %v", err)
	}

	rows, err := s.connPool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("execute a query failed: %v", err)
	}
	defer rows.Close()

	var questions []*Question

	for rows.Next() {
		question, err := readQuestion(rows)
		if err != nil {
			return nil, fmt.Errorf("read question failed: %v", question)
		}
		questions = append(questions, question)
	}

	if len(questions) == 1 {
		return questions[0], nil
	}

	return nil, fmt.Errorf("questions arr failed: %v", err)
}

func readQuestion(row pgx.Row) (*Question, error) {
	var q Question

	err := row.Scan(&q.Id, &q.Text, &q.SymptomId)
	if err != nil {
		return nil, err
	}

	return &q, nil
}
