package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

// Simple utility to check if all image files in CSV exist
func checkImages() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run check_images.go <csv_file>")
	}

	csvPath := os.Args[1]

	file, err := os.Open(csvPath)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
	}

	if len(records) == 0 {
		log.Fatal("CSV file is empty")
	}

	// Skip header if it exists
	startIdx := 0
	if len(records) > 0 {
		firstRow := strings.ToLower(strings.Join(records[0], "|"))
		if strings.Contains(firstRow, "file_path") || strings.Contains(firstRow, "title") {
			startIdx = 1
		}
	}

	allExist := true
	for i := startIdx; i < len(records); i++ {
		row := records[i]
		if len(row) == 0 || row[0] == "" {
			continue
		}

		filePath := row[0]
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("❌ Missing file: %s\n", filePath)
			allExist = false
		} else {
			fmt.Printf("✅ Found: %s\n", filePath)
		}
	}

	if !allExist {
		os.Exit(1)
	}

	fmt.Println("🎉 All image files found!")
}
