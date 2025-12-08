package service

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func (s *promptService) CreateMigrationFile(fileName string) error {

	fileName = strings.ToLower(fileName)

	dir := "migrations"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	timestamp := time.Now().Format("20060102_150405")

	upFile := fmt.Sprintf("%s/%s_%s.up.sql", dir, timestamp, fileName)
	if err := createEmptyFile(upFile); err != nil {
		return err
	}

	downFile := fmt.Sprintf("%s/%s_%s.down.sql", dir, timestamp, fileName)
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

	dir := "seeds"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	timestamp := time.Now().Format("20060102_150405")

	seedFile := fmt.Sprintf("%s/%s_%s.sql", dir, timestamp, fileName)
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
