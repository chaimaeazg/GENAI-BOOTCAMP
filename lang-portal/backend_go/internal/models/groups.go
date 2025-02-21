package models

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func GetAllGroups(db *sqlx.DB) ([]Group, error) {
	var groups []Group
	err := db.Select(&groups, "SELECT DISTINCT id, name FROM groups ORDER BY id LIMIT 10")
	return groups, err
}

func GetGroupByID(db *sqlx.DB, id int) (Group, error) {
	var group Group
	err := db.Get(&group, "SELECT id, name FROM groups WHERE id = $1", id)
	return group, err
}

func GetWordsByGroup(db *sqlx.DB, groupID int) ([]Word, error) {
	var words []Word
	query := `
		SELECT w.id, w.spanish, w.english
		FROM words w
		JOIN words_groups wg ON w.id = wg.word_id
		WHERE wg.group_id = $1`
	err := db.Select(&words, query, groupID)
	if err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}
	return words, nil
}

func GetStudySessionsByGroup(db *sqlx.DB, groupID int) ([]StudySession, error) {
	var sessions []StudySession
	err := db.Select(&sessions, `
		SELECT id, group_id 
		FROM study_sessions 
		WHERE group_id = $1`, groupID)
	return sessions, err
}

func GetOrCreateGroup(db sqlx.Ext, name string) (int, error) {
	var group Group
	err := sqlx.Get(db, &group, "SELECT id FROM groups WHERE name = ?", name)
	if err == nil {
		return group.ID, nil
	}

	if err == sql.ErrNoRows {
		res, err := sqlx.NamedExec(db,
			"INSERT INTO groups (name) VALUES (:name)",
			map[string]interface{}{"name": name},
		)
		if err != nil {
			return 0, fmt.Errorf("failed to insert group: %w", err)
		}

		id, err := res.LastInsertId()
		return int(id), err
	}

	return 0, fmt.Errorf("failed to query group: %w", err)
}
