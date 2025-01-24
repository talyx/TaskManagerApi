package models

import "time"

type Project struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	UserID      uint      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`

	User User `gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
}
