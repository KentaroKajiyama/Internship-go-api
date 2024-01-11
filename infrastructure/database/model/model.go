package model

import (
	"time"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ToDo struct {
	ID          string `gorm:"primaryKey"`
	ToDoID      string `gorm:"primaryKey"`
	Title       string
	Description string
	IsDeletable bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tag struct {
	ID        string `gorm:"primaryKey"`
	TagID     string `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
