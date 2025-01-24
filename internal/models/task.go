package models

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	ProjectID   uint      `gorm:"not null"`
	Status      string    `gorm:"type:varchar(20);default:'pending'"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`

	Project Project `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
}
