package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

func RecordReviewItem(db sqlx.Ext, sessionID int, wordID int, correct bool) error {
	_, err := sqlx.NamedExec(db, `
		INSERT INTO word_review_items 
		(study_session_id, word_id, correct, created_at)
		VALUES (:study_session_id, :word_id, :correct, :created_at)`,
		map[string]interface{}{
			"study_session_id": sessionID,
			"word_id":          wordID,
			"correct":          correct,
			"created_at":       time.Now().UTC(),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to insert review item: %w", err)
	}
	return nil
}
