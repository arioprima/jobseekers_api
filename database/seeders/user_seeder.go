package seeders

import (
	"github.com/arioprima/jobseekers_api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err) // Handle error appropriately, panic here for simplicity
	}
	return string(hashedPassword)
}

func UserSeeder(db *gorm.DB) {
	user := []models.User{
		{
			ID:         "019088d9-2143-7f1e-9dd3-1c696dd8aa25",
			BiodataId:  "019088d9-2143-7a83-a0be-2c2e0a5fecfc",
			Password:   HashPassword("test1234"),
			IsActive:   true,
			IsVerified: true,
			RoleId:     "019047ca-f542-7182-8b6b-7978f905dfe7",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	for _, u := range user {
		result := db.Create(&u)
		if result.Error != nil {
			log.Fatalf("Failed to seed user data: %v", result.Error)
		}
	}
}
