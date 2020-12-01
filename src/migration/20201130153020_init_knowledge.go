package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upInitKnowledge, downInitKnowledge)
}

func upInitKnowledge(tx *sql.Tx) error {
	_, err := tx.Exec(`
INSERT INTO knowledge (id_symptom, is_root, id_question, id_true_question, id_false_question, id_diagnosis) 
VALUES 
       (1, TRUE, 4, 29, 30, NULL),
       (2, TRUE, 1, 2, 3, NULL),
       (3, TRUE, 1, 2, 3, NULL),
       (4, TRUE, 1, 2, 3, NULL),
       (5, TRUE, 1, 2, 3, NULL),
       (6, TRUE, 1, 2, 3, NULL),
       (7, TRUE, 1, 2, 3, NULL),
       (8, TRUE, 1, 2, 3, NULL),
       (9, TRUE, 1, 2, 3, NULL),
       (10, TRUE, 1, 2, 3, NULL),
       (11, TRUE, 1, 2, 3, NULL),
       (12, TRUE, 1, 2, 3, NULL),
       (13, TRUE, 1, 2, 3, NULL),
       (14, TRUE, 1, 2, 3, NULL),
       (15, TRUE, 1, 2, 3, NULL),
       (16, TRUE, 1, 2, 3, NULL),
       (17, TRUE, 1, 2, 3, NULL),
       (18, TRUE, 1, 2, 3, NULL),
       (19, TRUE, 1, 2, 3, NULL),
       (20, TRUE, 1, 2, 3, NULL),
       (21, TRUE, 1, 2, 3, NULL),
       (22, TRUE, 1, 2, 3, NULL),
       (23, TRUE, 1, 2, 3, NULL),
       (24, TRUE, 1, 2, 3, NULL),
       (25, TRUE, 1, 2, 3, NULL),
       (26, TRUE, 1, 2, 3, NULL),
       (27, TRUE, 1, 2, 3, NULL),
       (28, TRUE, 1, 2, 3, NULL),
       (1, FALSE, NULL, NULL, NULL, 5), 
       (1, FALSE, 3, 31, 32, NULL),
       (1, FALSE, 2, 33, 34, NULL),
       (1, FALSE, 1, 35, 36, NULL),
       (1, FALSE, NULL, NULL, NULL, 4),
       (1, FALSE, NULL, NULL, NULL, 3),
       (1, FALSE, NULL, NULL, NULL, 2),
       (1, FALSE, NULL, NULL, NULL, 1);

`)
	//36
	return err
}

func downInitKnowledge(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
