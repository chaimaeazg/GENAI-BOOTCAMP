package models

import "github.com/jmoiron/sqlx"

type QuickStats struct {
	SuccessRate        float64 `json:"success_rate"`
	TotalStudySessions int     `json:"total_study_sessions"`
	TotalActiveGroups  int     `json:"total_active_groups"`
	StudyStreakDays    int     `json:"study_streak_days"`
}

func CalculateQuickStats(db *sqlx.DB) (QuickStats, error) {
	var stats QuickStats

	// Success Rate - handle case of no reviews
	err := db.Get(&stats.SuccessRate, `
		SELECT COALESCE(
			(SUM(CASE WHEN correct THEN 1.0 ELSE 0.0 END) * 100.0) / 
			NULLIF(COUNT(*), 0), 
			0.0
		) FROM word_review_items`)
	if err != nil {
		stats.SuccessRate = 0.0
	}

	// Total Study Sessions
	err = db.Get(&stats.TotalStudySessions, `SELECT COUNT(*) FROM study_sessions`)
	if err != nil {
		stats.TotalStudySessions = 0
	}

	// Active Groups
	err = db.Get(&stats.TotalActiveGroups, `
		SELECT COUNT(DISTINCT group_id) FROM study_sessions`)
	if err != nil {
		stats.TotalActiveGroups = 0
	}

	return stats, nil
}
