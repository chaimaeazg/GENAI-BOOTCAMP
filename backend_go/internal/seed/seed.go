package seed

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"backend_go/internal/models" // Changed from lang-portal

	"github.com/jmoiron/sqlx"
)

func ProcessSeedFile(db *sqlx.DB, filePath string, groupName string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read seed file: %v", err)
	}

	var importData struct {
		Words []models.Word `json:"words"`
	}

	if err := json.Unmarshal(data, &importData); err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	tx := db.MustBegin()
	defer tx.Rollback()

	// Create group if not exists
	groupID, err := models.GetOrCreateGroup(tx, groupName)
	if err != nil {
		return fmt.Errorf("failed to create group: %v", err)
	}

	// Insert words and create group associations
	for _, word := range importData.Words {
		res, err := tx.NamedExec(
			`INSERT INTO words (spanish, english) 
			VALUES (:spanish, :english)`,
			word,
		)
		if err != nil {
			return fmt.Errorf("failed to insert word: %v", err)
		}

		wordID, _ := res.LastInsertId()
		_, err = tx.Exec(
			`INSERT INTO words_groups (word_id, group_id) 
			VALUES (?, ?)`,
			wordID, groupID,
		)
		if err != nil {
			return fmt.Errorf("failed to associate word with group: %v", err)
		}
	}

	return tx.Commit()
}

// Add helper function to get data directory
func GetDataDir() string {
	return filepath.Join("internal", "seed", "data")
}
