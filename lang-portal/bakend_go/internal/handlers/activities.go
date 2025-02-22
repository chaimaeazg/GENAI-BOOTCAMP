package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"backend_go/internal/models"
)

func GetStudyActivity(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	activity, err := models.GetStudyActivityByID(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func GetActivitySuccessRate(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	activityID, _ := strconv.Atoi(c.Param("id"))
	successRate, err := models.CalculateActivitySuccessRate(db, activityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate success rate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"activity_id":  activityID,
		"success_rate": successRate,
	})
}

func GetAllStudyActivities(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	activities, err := models.GetAllStudyActivities(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve activities"})
		return
	}

	c.JSON(http.StatusOK, activities)
}

func CreateStudyActivity(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	var activity models.StudyActivity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := models.CreateStudyActivity(db, activity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create activity"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":      id,
		"name":    activity.Name,
		"message": "Activity created successfully",
	})
}

func GetActivityStudySessions(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	activityID, _ := strconv.Atoi(c.Param("id"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	sessions, total, err := models.GetActivitySessions(db, activityID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sessions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": sessions,
		"pagination": gin.H{
			"current_page":   page,
			"total_pages":    (total + limit - 1) / limit,
			"total_items":    total,
			"items_per_page": limit,
		},
	})
}

func LaunchStudyActivity(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	activityID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	launchURL, err := models.GetActivityLaunchURL(db, activityID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Activity launched successfully",
		"launch_url": launchURL,
	})
}
