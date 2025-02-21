package models

import (
	"github.com/jmoiron/sqlx"
)

func ResetStudyHistory(db *sqlx.DB) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM study_sessions`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DELETE FROM word_review_items`)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func FullSystemReset(db *sqlx.DB) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	tables := []string{
		"words", "groups", "words_groups",
		"study_sessions", "word_review_items",
	}

	for _, table := range tables {
		_, err = tx.Exec(`DELETE FROM ` + table)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
