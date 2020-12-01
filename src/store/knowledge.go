package store

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v4"
)

type Knowledge struct {
	Id              int
	SymptomId       int
	IsRoot          bool
	QuestionId      *int
	TrueQuestionId  *int
	FalseQuestionId *int
	DiagnosisId     *int
}

func (s *Store) CreateKnowledge(ctx context.Context, knowledge Knowledge) error {
	sql, _, err := goqu.Insert("knowledge").
		Rows(goqu.Record{
			"id_symptom":        knowledge.SymptomId,
			"is_root":           knowledge.IsRoot,
			"id_question":       knowledge.QuestionId,
			"id_true_question":  knowledge.TrueQuestionId,
			"id_false_question": knowledge.FalseQuestionId,
			"id_diagnosis":      knowledge.DiagnosisId,
		}).ToSQL()
	if err != nil {
		return fmt.Errorf("sql query build failed: %v", err)
	}
	if _, err := s.connPool.Exec(ctx, sql); err != nil {
		return fmt.Errorf("execute a query failed: %v", err)
	}

	return nil
}

func (s *Store) GetAllKnowledge(ctx context.Context) ([]*Knowledge, error) {
	sql, _, err := goqu.Select().From("knowledge").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("sql query build failed: %v", err)
	}

	rows, err := s.connPool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("execute a query failed: %v", err)
	}
	defer rows.Close()

	var knowledgeArr []*Knowledge

	for rows.Next() {
		knowledge, err := readKnowledge(rows)
		if err != nil {
			return nil, fmt.Errorf("read question failed: %v", knowledge)
		}
		knowledgeArr = append(knowledgeArr, knowledge)
	}

	return knowledgeArr, nil
}

func (s *Store) GetKnowledgeById(ctx context.Context, id int) (*Knowledge, error) {
	sql, _, err := goqu.Select().From("knowledge").Where(goqu.C("id").Eq(id)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("sql query build failed: %v", err)
	}

	rows, err := s.connPool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("execute a query failed: %v", err)
	}
	defer rows.Close()

	var knowledgeArr []*Knowledge

	for rows.Next() {
		knowledge, err := readKnowledge(rows)
		if err != nil {
			return nil, fmt.Errorf("read question failed: %v", knowledge)
		}
		knowledgeArr = append(knowledgeArr, knowledge)
	}

	if len(knowledgeArr) == 1 {
		return knowledgeArr[0], nil
	}

	return nil, fmt.Errorf("knowledge arr failed: %v", err)
}

func (s *Store) GetKnowledgeBySymptomId(ctx context.Context, symptomId int) ([]*Knowledge, error) {
	sql, _, err := goqu.Select().
		From("knowledge").
		Where(goqu.C("id_symptom").Eq(symptomId)).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("sql query build failed: %v", err)
	}

	rows, err := s.connPool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("execute a query failed: %v", err)
	}
	defer rows.Close()

	var knowledgeArr []*Knowledge

	for rows.Next() {
		knowledge, err := readKnowledge(rows)
		if err != nil {
			return nil, fmt.Errorf("read question failed: %v", knowledge)
		}
		knowledgeArr = append(knowledgeArr, knowledge)
	}

	return knowledgeArr, nil
}

func (s *Store) GetRootKnowledge(ctx context.Context, symptomId int) (*Knowledge, error) {
	sql, _, err := goqu.Select().From("knowledge").
		Where(goqu.C("id_symptom").Eq(symptomId),
			goqu.C("is_root").Eq(true)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("sql query build failed: %v", err)
	}

	rows, err := s.connPool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("execute a query failed: %v", err)
	}
	defer rows.Close()

	var knowledgeArr []*Knowledge

	for rows.Next() {
		knowledge, err := readKnowledge(rows)
		if err != nil {
			return nil, fmt.Errorf("read question failed: %v", knowledge)
		}
		knowledgeArr = append(knowledgeArr, knowledge)
	}

	if len(knowledgeArr) == 1 {
		return knowledgeArr[0], nil
	}

	return nil, fmt.Errorf("knowledge arr failed: %v", err)
}

func readKnowledge(row pgx.Row) (*Knowledge, error) {
	var d Knowledge

	err := row.Scan(&d.Id, &d.SymptomId, &d.IsRoot, &d.QuestionId, &d.TrueQuestionId, &d.FalseQuestionId, &d.DiagnosisId)
	if err != nil {
		return nil, err
	}

	return &d, nil
}