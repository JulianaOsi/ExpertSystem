package store

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v4"
)

type Diagnosis struct {
	Id        int
	Name      string
	SymptomId int
}

func (s *Store) CreateOrUpdateDiagnosis(ctx context.Context, diagnosis *Diagnosis) error {
	sql, _, err := goqu.Insert("diagnosis").
		Rows(goqu.Record{
			"id":         diagnosis.Id,
			"name":       diagnosis.Name,
			"id_symptom": diagnosis.SymptomId,
		}).
		OnConflict(goqu.DoUpdate("id", goqu.Record{
			"name":       diagnosis.Name,
			"id_symptom": diagnosis.SymptomId,
		})).ToSQL()
	if err != nil {
		return fmt.Errorf("sql query build failed: %v", err)
	}
	if _, err := s.connPool.Exec(ctx, sql); err != nil {
		return fmt.Errorf("execute a query failed: %v", err)
	}

	return nil
}

func (s *Store) GetAllDiagnoses(ctx context.Context) ([]*Diagnosis, error) {
	sql, _, err := goqu.Select().From("diagnosis").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("sql query build failed: %v", err)
	}

	rows, err := s.connPool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("execute a query failed: %v", err)
	}
	defer rows.Close()

	var diagnoses []*Diagnosis

	for rows.Next() {
		diagnosis, err := readDiagnosis(rows)
		if err != nil {
			return nil, fmt.Errorf("read question failed: %v", diagnosis)
		}
		diagnoses = append(diagnoses, diagnosis)
	}

	return diagnoses, nil
}

func (s *Store) GetDiagnosisById(ctx context.Context, id int) (*Diagnosis, error) {
	sql, _, err := goqu.Select().
		From("diagnosis").
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

	var diagnosis []*Diagnosis

	for rows.Next() {
		d, err := readDiagnosis(rows)
		if err != nil {
			return nil, fmt.Errorf("read diagnosis failed: %v", d)
		}
		diagnosis = append(diagnosis, d)
	}

	if len(diagnosis) == 1 {
		return diagnosis[0], nil
	}

	return nil, fmt.Errorf("diagnosis arr failed: %v", err)
}

func (s *Store) GetDiagnosisBySymptomId(ctx context.Context, symptomId int) ([]*Diagnosis, error) {
	sql, _, err := goqu.Select().
		From("diagnosis").
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

	var diagnosis []*Diagnosis

	for rows.Next() {
		d, err := readDiagnosis(rows)
		if err != nil {
			return nil, fmt.Errorf("read diagnosis failed: %v", d)
		}
		diagnosis = append(diagnosis, d)
	}

	return diagnosis, nil
}

func readDiagnosis(row pgx.Row) (*Diagnosis, error) {
	var d Diagnosis

	err := row.Scan(&d.Id, &d.Name, &d.SymptomId)
	if err != nil {
		return nil, err
	}

	return &d, nil
}
