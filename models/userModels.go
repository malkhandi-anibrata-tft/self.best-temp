package models

import "time"

type UserDetails struct {
	Id        uint64 `gorm:"primaryKey;autoIncrement"`
	FirstName string
	LastName  string
	Email     string `gorm:"unique;not null"`
	Password  string
	CreatedAt time.Time
}
