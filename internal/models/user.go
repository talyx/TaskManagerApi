package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Names        string    `gorm:"type:varchar(100);not null"`
	Email        string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
