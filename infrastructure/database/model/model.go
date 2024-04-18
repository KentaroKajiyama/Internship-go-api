package model

import (
	"time"
)

type User struct {
	Id          string `gorm:"type:uuid;primaryKey"`
	FirebaseUid string `gorm:"uniqueIndex;not null"`
	Name        string
	Email       string `gorm:"uniqueIndex;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Todo struct {
	Id          string `gorm:"primaryKey"`
	TodoId      string `gorm:"primaryKey"`
	Title       string
	Description string
	IsDeletable bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tag struct {
	Id        string `gorm:"primaryKey"`
	TagId     uint64 `gorm:"primaryKey;autoincrement"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TodoTag struct {
	TodoId string `gorm:"primaryKey"`
	TagId  uint64 `gorm:"primaryKey"`
}
