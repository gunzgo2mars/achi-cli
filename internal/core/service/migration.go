package service

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gunzgo2mars/achi-cli/internal/constants"
)

func (s *promptService) CreateMigrationFile(fileName string) error {

	fileName = strings.ToLower(fileName)

	if err := os.MkdirAll(constants.MIGRATION_DIR, os.ModePerm); err != nil {
		return err
	}

	timestamp := time.Now().Format("20060102150405")

	upFile := fmt.Sprintf("%s/%s_%s.up.sql", constants.MIGRATION_DIR, timestamp, fileName)
	if err := createEmptyFile(upFile); err != nil {
		return err
	}

	downFile := fmt.Sprintf("%s/%s_%s.down.sql", constants.MIGRATION_DIR, timestamp, fileName)
	if err := createEmptyFile(downFile); err != nil {
		return err
	}

	fmt.Println("Created migration:")
	fmt.Println(" →", upFile)
	fmt.Println(" →", downFile)

	return nil

}

func (s *promptService) CreateSeederFile(fileName string) error {

	fileName = strings.ToLower(fileName)

	if err := os.MkdirAll(constants.SEEDER_DIR, os.ModePerm); err != nil {
		return err
	}

	timestamp := time.Now().Format("20060102150405")

	seedFile := fmt.Sprintf("%s/%s_%s.sql", constants.SEEDER_DIR, timestamp, fileName)
	if err := createEmptyFile(seedFile); err != nil {
		return err
	}

	fmt.Println("Created seeder:")
	fmt.Println(" →", seedFile)

	return nil

}

func createEmptyFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write header comment
	_, err = f.WriteString("-- Write SQL here\n")
	if err != nil {
		return err
	}

	return nil

}
