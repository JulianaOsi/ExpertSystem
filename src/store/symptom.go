package store

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v4"
)

type Symptom struct {
	Id   int
	Name string
}

func (s *Store) CreateOrUpdateSymptom(ctx context.Context, symptom *Symptom) error {
	sql, _, err := goqu.Insert("symptom").
		Rows(goqu.Record{
			"id":   symptom.Id,
			"name": symptom.Name,
		}).
		OnConflict(goqu.DoUpdate("id", goqu.Record{
			"name": symptom.Name,
		})).ToSQL()
	if err != nil {
		return fmt.Errorf("sql query build failed: %v", err)
	}
	if _, err := s.connPool.Exec(ctx, sql); err != nil {
		return fmt.Errorf("execute a query failed: %v", err)
	}

	return nil
}

func (s *Store) GetAllSymptoms(ctx context.Context) ([]*Symptom, error) {
	sql, _, err := goqu.Select().From("symptom").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("sql query build failed: %v", err)
	}

	rows, err := s.connPool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("execute a query failed: %v", err)
	}
	defer rows.Close()

	var symptoms []*Symptom

	for rows.Next() {
		symptom, err := readSymptom(rows)
		if err != nil {
			return nil, fmt.Errorf("read symptom failed: %v", symptom)
		}
		symptoms = append(symptoms, symptom)
	}

	return symptoms, nil
}

func readSymptom(row pgx.Row) (*Symptom, error) {
	var s Symptom

	err := row.Scan(&s.Id, &s.Name)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
