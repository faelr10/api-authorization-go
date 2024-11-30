package models

import "time"

type User struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string `gorm:"not null"`
    CreatedAt time.Time
}
