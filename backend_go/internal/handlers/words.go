package handlers

import (
	"net/http"
	"strconv"

	"backend_go/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetWordDetails(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	wordID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid word ID"})
		return
	}

	word, err := models.GetWordByID(db, wordID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Word not found"})
		return
	}

	stats, err := models.GetWordStats(db, wordID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get word stats"})
		return
	}

	groups, err := models.GetWordGroups(db, wordID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get word groups"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"spanish": word.Spanish,
		"english": word.English,
		"stats":   stats,
		"groups":  groups,
	})
}

func GetWordStudySessions(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	wordID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid word ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	sessions, total, err := models.GetWordStudySessions(db, wordID, page, limit)
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

func GetWords(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	words, total, err := models.GetPaginatedWords(db, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve words"})
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
