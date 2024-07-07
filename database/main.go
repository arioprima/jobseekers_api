package main

import (
	"fmt"
	"github.com/arioprima/jobseekers_api/config"
	"github.com/arioprima/jobseekers_api/database/seeders"
	"log"
	"os"
)

func main() {
	configPath := "."
	// Load configuration
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Open database connection
	db, err := config.OpenConnection(&cfg)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	log.Println("Database connection opened")
	// Seed data
	seeders.SeedRole(db)
	seeders.SeedBio(db)
	seeders.UserSeeder(db)

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting DB object: %v", err)
	}
	defer sqlDB.Close()

	// Selesai
	fmt.Println("Successfully inserted data")

	// Exit program
	os.Exit(0)
}
