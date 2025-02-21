package handlers

import (
	"net/http"
	"strconv"

	"backend_go/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RecordReview(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	sessionID, _ := strconv.Atoi(c.Param("id"))
	wordID, _ := strconv.Atoi(c.Param("word_id"))

	var review struct{ Correct bool }
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := models.RecordReviewItem(db, sessionID, wordID, review.Correct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record review"})
		return
	}

	c.Status(http.StatusNoContent)
}
