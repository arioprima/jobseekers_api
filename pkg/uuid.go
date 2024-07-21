package pkg

import "github.com/google/uuid"

func GenerateUUID() string {
	generateUuid, err := uuid.NewV7()
	if err != nil {
		return ""
	}
	return generateUuid.String()
}
