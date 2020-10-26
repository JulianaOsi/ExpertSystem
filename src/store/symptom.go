package store

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
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
