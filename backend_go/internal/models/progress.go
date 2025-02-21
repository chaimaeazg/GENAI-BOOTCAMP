package models

import "github.com/jmoiron/sqlx"

type Progress struct {
	Studied int
	Total   int
}

func CalculateStudyProgress(db *sqlx.DB) (Progress, error) {
	var p Progress
	err := db.Get(&p.Studied, `
        SELECT COUNT(DISTINCT word_id) 
        FROM word_review_items
        WHERE correct = true
    `)
	if err != nil {
		return p, err
	}

	err = db.Get(&p.Total, "SELECT COUNT(*) FROM words")
	return p, err
}
