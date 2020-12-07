package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upInitSchema, downInitSchema)
}

func upInitSchema(tx *sql.Tx) error {
	_, err := tx.Exec(`
CREATE TABLE IF NOT EXISTS diagnosis
(
    id   		INT 	NOT NULL,
    title   	TEXT    NOT NULL,
	id_symptom 	INT		NOT NULL
);

CREATE TABLE IF NOT EXISTS knowledge
(
    id 					SERIAL PRIMARY KEY,
    id_inner			INT NOT NULL,
    id_symptom          INT NOT NULL,
    id_question         INT,
    id_true_question    INT,
    id_false_question 	INT,
    id_diagnosis		INT
);

CREATE TABLE IF NOT EXISTS question
(
    id       	INT		NOT NULL,
    text     	TEXT	NOT NULL,
    symptom_id	INT 	NOT NULL
);

CREATE TABLE IF NOT EXISTS specialty
(
    id          	INT		NOT NULL,
    title  			TEXT	NOT NULL,
    diagnosis_id	INT
);

CREATE TABLE IF NOT EXISTS symptom
(
    id          	INT		NOT NULL,
    title  			TEXT	NOT NULL
);
`)
	return err
}

func downInitSchema(tx *sql.Tx) error {
	_, err := tx.Exec(`
DROP TABLE diagnosis CASCADE;
DROP TABLE knowledge CASCADE;
DROP TABLE question CASCADE;
DROP TABLE specialty CASCADE;
DROP TABLE symptom CASCADE;

`)
	return err
}
