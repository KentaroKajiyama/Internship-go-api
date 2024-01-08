package uuid

import (
	"github.com/google/uuid"
)

func NewUUID() string {
	return uuid.NewString()
}

func IsValid(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}
