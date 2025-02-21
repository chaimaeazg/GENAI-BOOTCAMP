package handlers

import (
	"backend_go/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func CreateStudySession(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	var session struct {
		GroupID int `json:"group_id"`
	}
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	sessionID, err := models.CreateStudySession(db, session.GroupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         sessionID,
		"group_id":   session.GroupID,
		"created_at": time.Now().UTC(),
	})
}

func GetAllStudySessions(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	sessions, total, err := models.GetPaginatedStudySessions(db, page, limit)
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

func GetStudySessionDetails(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
		return
	}

	session, err := models.GetStudySessionByID(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	c.JSON(http.StatusOK, session)
}

func GetSessionWords(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	sessionID, _ := strconv.Atoi(c.Param("id"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	words, total, err := models.GetSessionWords(db, sessionID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session words"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": words,
		"pagination": gin.H{
			"current_page":   page,
			"total_pages":    (total + limit - 1) / limit,
			"total_items":    total,
			"items_per_page": limit,
		},
	})
}
