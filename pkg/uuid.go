package pkg

import "github.com/google/uuid"

func NewUuid() string {
	id := uuid.New()
	return id.String()
}
