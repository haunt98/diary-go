package main

import (
	"cmp"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	locationGMT07, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		log.Fatalln("Failed to load GMT+07 timezone")
	}

	today := time.Now().In(locationGMT07)

	// Obsidian daily note format
	filename := filepath.Join(
		cmp.Or(os.Getenv("DIARY_DIR"), "."),
		today.Format("2006-01-02")+".md",
	)

	// Create file if not exist
	// See os.Create
	f, err := os.OpenFile(filename, os.O_CREATE, 0o666)
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", filename, err)
	}
	f.Close()
}
