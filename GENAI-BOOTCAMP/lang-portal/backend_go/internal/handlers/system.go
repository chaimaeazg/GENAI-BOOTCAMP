package handlers

import (
	"net/http"

	"backend_go/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetQuickStats(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	stats, err := models.CalculateQuickStats(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate stats"})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func GetLastStudySession(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	session, err := models.GetLatestStudySession(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, session)
}

func ResetHistory(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	err := models.ResetStudyHistory(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Study history reset successfully",
	})
}

func FullReset(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	err := models.FullSystemReset(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset system"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "System reset and reseeded successfully",
	})
}
