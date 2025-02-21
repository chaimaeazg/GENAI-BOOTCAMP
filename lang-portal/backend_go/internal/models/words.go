package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func InsertWord(db sqlx.Ext, word Word) (int, error) {
	res, err := sqlx.NamedExec(db, `
		INSERT INTO words (spanish, english, parts)
		VALUES (:spanish, :english, :parts)`,
		word,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to insert word: %w", err)
	}

	id, err := res.LastInsertId()
	return int(id), err
}

func GetWordByID(db *sqlx.DB, id int) (Word, error) {
	var word Word
	err := db.Get(&word, "SELECT * FROM words WHERE id = ?", id)
	return word, err
}

func GetWordStats(db *sqlx.DB, wordID int) (map[string]int, error) {
	stats := make(map[string]interface{})
	err := db.QueryRowx(`
		SELECT 
			COUNT(CASE WHEN correct = true THEN 1 END) as correct_count,
			COUNT(CASE WHEN correct = false THEN 1 END) as wrong_count
		FROM word_review_items
		WHERE word_id = $1
	`, wordID).MapScan(stats)

	result := make(map[string]int)
	for k, v := range stats {
		if num, ok := v.(int64); ok {
			result[k] = int(num)
		}
	}
	return result, err
}

func GetWordGroups(db *sqlx.DB, wordID int) ([]Group, error) {
	var groups []Group
	err := db.Select(&groups, `
		SELECT g.id, g.name 
		FROM groups g
		JOIN words_groups wg ON g.id = wg.group_id
		WHERE wg.word_id = $1
	`, wordID)
	return groups, err
}

func GetWordStudySessions(db *sqlx.DB, wordID int, page int, limit int) ([]StudySession, int, error) {
	var sessions []StudySession
	offset := (page - 1) * limit

	// Get paginated sessions
	err := db.Select(&sessions, `
		SELECT s.id, s.group_id, s.created_at 
		FROM study_sessions s
		JOIN word_review_items wr ON s.id = wr.study_session_id
		WHERE wr.word_id = $1
		LIMIT $2 OFFSET $3
	`, wordID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	// Get total count
	var total int
	err = db.Get(&total, `
		SELECT COUNT(DISTINCT s.id)
		FROM study_sessions s
		JOIN word_review_items wr ON s.id = wr.study_session_id
		WHERE wr.word_id = $1
	`, wordID)

	return sessions, total, err
}

func GetPaginatedWords(db *sqlx.DB, page int, limit int) ([]Word, int, error) {
	var words []Word
	offset := (page - 1) * limit

	err := db.Select(&words, `
		SELECT * FROM words 
		LIMIT $1 OFFSET $2`,
		limit, offset)
	if err != nil {
		return nil, 0, err
	}

	var total int
	err = db.Get(&total, "SELECT COUNT(*) FROM words")

	return words, total, err
}

func GetAllWords(db *sqlx.DB) ([]Word, error) {
	words := []Word{}
	err := db.Select(&words, "SELECT * FROM words")
	if err != nil {
		return nil, fmt.Errorf("failed to select words: %w", err)
	}
	return words, nil
}
