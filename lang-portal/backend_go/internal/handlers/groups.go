package handlers

import (
	"net/http"
	"strconv"

	"backend_go/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetAllGroups(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	groups, err := models.GetAllGroups(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve groups"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": groups,
		"pagination": gin.H{
			"current_page":   1,
			"total_pages":    1,
			"items_per_page": 10,
		},
	})
}

func GetGroupByID(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	group, err := models.GetGroupByID(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}

	c.JSON(http.StatusOK, group)
}

func GetGroupWords(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	words, err := models.GetWordsByGroup(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve group words"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": words,
		"pagination": gin.H{
			"current_page":   1,
			"total_pages":    1,
			"items_per_page": len(words),
		},
	})
}

func GetGroupStudySessions(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	sessions, err := models.GetStudySessionsByGroup(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve study sessions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": sessions})
}
