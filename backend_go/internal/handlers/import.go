package handlers

import (
	"backend_go/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type ImportData struct {
	Words []struct {
		ID      int    `json:"id"`
		Spanish string `json:"spanish"`
		English string `json:"english"`
		GroupID int    `json:"group_id"`
	} `json:"words"`
	Groups []models.Group `json:"groups"`
}

func ImportInitialData(c *gin.Context) {
	db := c.MustGet("db").(*sqlx.DB)

	filePath := "internal/seed/data/initial_data.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Import file not found: %v", err)})
		return
	}

	var importData ImportData
	if err := json.Unmarshal(data, &importData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid JSON format: %v", err)})
		return
	}

	tx := db.MustBegin()

	// Insert groups
	for _, group := range importData.Groups {
		_, err := tx.NamedExec(
			`INSERT INTO groups (id, name) VALUES (:id, :name)`,
			group,
		)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Group import failed: %v", err)})
			return
		}
	}

	// Insert words and group relationships
	for _, word := range importData.Words {
		_, err := tx.NamedExec(
			`INSERT INTO words (id, spanish, english)
			VALUES (:id, :spanish, :english)`,
			map[string]interface{}{
				"id":      word.ID,
				"spanish": word.Spanish,
				"english": word.English,
			},
		)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Word import failed: %v", err)})
			return
		}

		_, err = tx.Exec(
			`INSERT INTO words_groups (word_id, group_id) 
			VALUES (?, ?)`,
			word.ID, word.GroupID,
		)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Word group relationship import failed"})
			return
		}
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction commit failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data imported successfully",
		"stats": gin.H{
			"words":  len(importData.Words),
			"groups": len(importData.Groups),
		},
	})
}
