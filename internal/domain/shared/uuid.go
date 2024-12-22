package shared

import "github.com/google/uuid"

func NewId() *string {
	id := uuid.New().String()
	return &id
}