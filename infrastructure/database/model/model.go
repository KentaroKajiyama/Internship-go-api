package model

import (
	"time"
)

type User struct {
	Id        string `gorm:"primaryKey"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Todo struct {
	Id          string `gorm:"primaryKey"`
	TodoId      string `gorm:"primaryKey"`
	Title       string
	Text        string
	IsDeletable bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tag struct {
	Id        string `gorm:"primaryKey"`
	TagId     string `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
