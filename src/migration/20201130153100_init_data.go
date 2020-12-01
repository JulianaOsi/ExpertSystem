package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upInitData, downInitData)
}

func upInitData(tx *sql.Tx) error {
	_, err := tx.Exec(`
INSERT INTO specialty (id, title, diagnosis_id) 
VALUES 
       (1, 'врач-невролог', 11),
	   (2, 'врач-инфекционист', 9),
       (3, 'врач-терапевт', 5),
       (4, 'врач-пульмонолог', 5),
       (5, 'врач-хирург', 7),
       (6, 'врач-гастроэнтеролог', 8),
       (7, 'врач-травматолог', 1),
       (8, 'врач-кардиолог', 3),
       (9, 'врач-терапевт', 3);
`)
	return err
}

func downInitData(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
