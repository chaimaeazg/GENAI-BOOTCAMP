package models

import (
	"time"
)

type Word struct {
	ID      int    `json:"id" db:"id"`
	Spanish string `json:"spanish" db:"spanish"`
	English string `json:"english" db:"english"`
}

type WordGroup struct {
	ID      int `json:"id" db:"id"`
	WordID  int `json:"word_id" db:"word_id"`
	GroupID int `json:"group_id" db:"group_id"`
}

type Group struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type StudySession struct {
	ID      int `json:"id" db:"id"`
	GroupID int `json:"group_id" db:"group_id"`
}

type WordReviewItem struct {
	WordID         int    `json:"word_id" db:"word_id"`
	StudySessionID int    `json:"study_session_id" db:"study_session_id"`
	Correct        bool   `json:"correct" db:"correct"`
	CreatedAt      string `json:"created_at" db:"created_at"`
}

type LastSessionView struct {
	ID           int       `db:"id" json:"id"`
	GroupID      int       `db:"group_id" json:"group_id"`
	GroupName    string    `db:"group_name" json:"group_name"`
	ActivityName string    `db:"activity_name" json:"activity_type"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	Duration     int       `db:"duration" json:"duration"`
}
