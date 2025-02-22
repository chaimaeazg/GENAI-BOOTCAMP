//go:build mage
// +build mage

package main

import (
	"fmt"
	"log"
	"path/filepath"

	"backend_go/internal/models"
	"backend_go/internal/seed"

	"github.com/magefile/mage/mg"
)

type Seed mg.Namespace

// Seed database with test data
func (Seed) DB() error {
	db, err := models.InitDB("words.db")
	if err != nil {
		return fmt.Errorf("failed to initialize database: %v", err)
	}
	defer db.Close()

	seedFiles, err := filepath.Glob("db/seeds/test_data/*.json")
	if err != nil {
		return fmt.Errorf("failed to find seed files: %v", err)
	}

	for _, file := range seedFiles {
		groupName := filepath.Base(file)
		groupName = groupName[:len(groupName)-len(filepath.Ext(groupName))] // Remove extension

		log.Printf("Seeding group '%s' from %s", groupName, file)
		if err := seed.ProcessSeedFile(db, file, groupName); err != nil {
			return fmt.Errorf("failed to seed %s: %v", file, err)
		}
	}

	return nil
}
