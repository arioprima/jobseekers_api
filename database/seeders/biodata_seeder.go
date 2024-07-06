package seeders

import (
	"github.com/arioprima/jobseekers_api/models"
	"gorm.io/gorm"
	"log"
	"time"
)

func SeedBio(db *gorm.DB) {
	biodata := []models.Biodata{
		{
			ID:        "019088d9-2143-7a83-a0be-2c2e0a5fecfc",
			Firstname: "Tony",
			Lastname:  "Stark",
			Email:     "tonystark@gmail.com",
			Phone:     "08123456789",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, bio := range biodata {
		result := db.Create(&bio)
		if result.Error != nil {
			log.Fatalf("Failed to seed biodata data: %v", result.Error)
		}
	}
}
