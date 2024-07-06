package seeders

import (
	"github.com/arioprima/jobseekers_api/models"
	"gorm.io/gorm"
	"log"
	"time"
)

func SeedRole(db *gorm.DB) {
	roles := []models.UserRole{
		{
			ID:        "019047ca-f542-7182-8b6b-7978f905dfe7",
			Name:      "admin",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "019047ca-f542-71fe-9de6-c4919ed5c9ff",
			Name:      "company",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "019047ca-f542-71fe-9de6-c4919ed5c9ff",
			Name:      "user",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, role := range roles {
		result := db.Create(&role)
		if result.Error != nil {
			log.Fatalf("Failed to seed role data: %v", result.Error)
		}
	}
}
