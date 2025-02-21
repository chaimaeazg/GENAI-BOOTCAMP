package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type StudyActivity struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	LaunchURL   string `db:"launch_url"`
}

func GetStudyActivityByID(db *sqlx.DB, id int) (StudyActivity, error) {
	var activity StudyActivity
	err := db.Get(&activity, `
        SELECT id, name, description, launch_url 
        FROM study_activities 
        WHERE id = $1`, id)
	return activity, err
}

func CalculateActivitySuccessRate(db *sqlx.DB, activityID int) (float64, error) {
	var successRate float64
	err := db.Get(&successRate, `
        SELECT 
            (COUNT(CASE WHEN correct = true THEN 1 END) * 100.0) / COUNT(*)
        FROM word_review_items
        JOIN study_sessions ON word_review_items.study_session_id = study_sessions.id
        WHERE study_sessions.study_activity_id = $1
    `, activityID)
	return successRate, err
}

func GetAllStudyActivities(db *sqlx.DB) ([]StudyActivity, error) {
	var activities []StudyActivity
	err := db.Select(&activities, `
		SELECT id, name, description, launch_url 
		FROM study_activities`)
	return activities, err
}

func CreateStudyActivity(db sqlx.Ext, activity StudyActivity) (int, error) {
	res, err := sqlx.NamedExec(db, `
		INSERT INTO study_activities 
		(name, description, launch_url)
		VALUES (:name, :description, :launch_url)`,
		activity,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to insert activity: %w", err)
	}

	id, err := res.LastInsertId()
	return int(id), err
}

func GetActivitySessions(db *sqlx.DB, activityID int, page int, limit int) ([]StudySession, int, error) {
	var sessions []StudySession
	offset := (page - 1) * limit

	// Get sessions
	err := db.Select(&sessions, `
		SELECT id, group_id, created_at 
		FROM study_sessions 
		WHERE study_activity_id = $1
		LIMIT $2 OFFSET $3`,
		activityID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	// Get total count
	var total int
	err = db.Get(&total, `
		SELECT COUNT(*) 
		FROM study_sessions 
		WHERE study_activity_id = $1`,
		activityID)

	return sessions, total, err
}

func GetActivityLaunchURL(db *sqlx.DB, id int) (string, error) {
	var launchURL string
	err := db.Get(&launchURL, `
		SELECT launch_url 
		FROM study_activities 
		WHERE id = $1`, id)
	return launchURL, err
}

func CreateStudySession(db sqlx.Ext, groupID int) (int, error) {
	res, err := sqlx.NamedExec(db,
		`INSERT INTO study_sessions (group_id) VALUES (:group_id)`,
		map[string]interface{}{"group_id": groupID},
	)
	if err != nil {
		return 0, fmt.Errorf("failed to insert study session: %w", err)
	}

	id, err := res.LastInsertId()
	return int(id), err
}
