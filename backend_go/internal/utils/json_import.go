package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Word struct {
	Spanish string `json:"spanish"`
	English string `json:"english"`
}

func ImportWordsFromJSON(filePath string, db *sql.DB) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open JSON file: %w", err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var words []Word
	if err := json.Unmarshal(byteValue, &words); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	for _, word := range words {
		_, err := db.Exec("INSERT INTO words (spanish, english) VALUES (?, ?)", word.Spanish, word.English)
		if err != nil {
			return fmt.Errorf("failed to insert word into database: %w", err)
		}
	}

	return nil
}
