package store

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v4"
)

type Specialty struct {
	Id   int
	Name string
}

func (s *Store) CreateOrUpdateSpecialty(ctx context.Context, specialty *Specialty) error {
	sql, _, err := goqu.Insert("specialty").
		Rows(goqu.Record{
			"id":   specialty.Id,
			"name": specialty.Name,
		}).
		OnConflict(goqu.DoUpdate("id", goqu.Record{
			"name": specialty.Name,
		})).ToSQL()
	if err != nil {
		return fmt.Errorf("sql query build failed: %v", err)
	}
	if _, err := s.connPool.Exec(ctx, sql); err != nil {
		return fmt.Errorf("execute a query failed: %v", err)
	}

	return nil
}

func (s *Store) GetAllSpecialty(ctx context.Context) ([]*Specialty, error) {
	sql, _, err := goqu.Select().From("specialty").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("sql query build failed: %v", err)
	}

	rows, err := s.connPool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("execute a query failed: %v", err)
	}
	defer rows.Close()

	var specialty []*Specialty

	for rows.Next() {
		s, err := readSpecialty(rows)
		if err != nil {
			return nil, fmt.Errorf("read specialty failed: %v", s)
		}
		specialty = append(specialty, s)
	}

	return specialty, nil
}

func readSpecialty(row pgx.Row) (*Specialty, error) {
	var s Specialty

	err := row.Scan(&s.Id, &s.Name)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
