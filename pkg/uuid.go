package pkg

import "github.com/google/uuid"

func GenerateUUID() (string, error) {
	generateUuid, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return generateUuid.String(), nil
}
