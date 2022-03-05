package coreUuid

import "github.com/google/uuid"

// CreateUuid create uuid from google/uuid
func CreateUuid() string {
	return uuid.New().String()
}
