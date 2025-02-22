package models

import (
	"github.com/jmoiron/sqlx"
)

func GetPaginatedStudySessions(db *sqlx.DB, page int, limit int) ([]StudySession, int, error) {
	var sessions []StudySession
	offset := (page - 1) * limit

	// Get paginated sessions
	err := db.Select(&sessions, `
		SELECT id, group_id, created_at 
		FROM study_sessions 
		LIMIT $1 OFFSET $2`,
		limit, offset)
	if err != nil {
		return nil, 0, err
	}

	// Get total count
	var total int
	err = db.Get(&total, `SELECT COUNT(*) FROM study_sessions`)

	return sessions, total, err
}

func GetStudySessionByID(db *sqlx.DB, id int) (StudySession, error) {
	var session StudySession
	err := db.Get(&session, `
		SELECT id, group_id, created_at 
		FROM study_sessions 
		WHERE id = $1`, id)
	return session, err
}

func GetSessionWords(db *sqlx.DB, sessionID int, page int, limit int) ([]Word, int, error) {
	var words []Word
	offset := (page - 1) * limit

	// Get paginated words
	err := db.Select(&words, `
		SELECT w.* 
		FROM words w
		JOIN word_review_items wr ON w.id = wr.word_id
		WHERE wr.study_session_id = $1
		LIMIT $2 OFFSET $3`,
		sessionID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	// Get total count
	var total int
	err = db.Get(&total, `
		SELECT COUNT(*) 
		FROM word_review_items 
		WHERE study_session_id = $1`,
		sessionID)
	if err != nil {
		return nil, 0, err
	}

	return words, total, err
}

func GetLatestStudySession(db *sqlx.DB) (LastSessionView, error) {
	var session LastSessionView
	err := db.Get(&session, `
		SELECT 
			s.id,
			s.group_id,
			g.name AS group_name,
			a.name AS activity_name,
			s.created_at,
			CAST(JULIANDAY(s.ended_at) - JULIANDAY(s.created_at) * 24 * 60 AS INTEGER) AS duration
		FROM study_sessions s
		JOIN groups g ON s.group_id = g.id
		JOIN study_activities a ON s.study_activity_id = a.id
		ORDER BY s.created_at DESC
		LIMIT 1
	`)
	return session, err
}
