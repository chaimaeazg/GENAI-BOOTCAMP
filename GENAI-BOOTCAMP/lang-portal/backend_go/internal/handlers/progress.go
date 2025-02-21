package handlers

import (
	"backend_go/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetStudyProgress(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	progress, err := models.CalculateStudyProgress(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate progress"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_words_studied":   progress.Studied,
		"total_available_words": progress.Total,
	})
}
