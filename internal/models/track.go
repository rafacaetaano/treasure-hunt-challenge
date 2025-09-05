package models

import "time"

type Track struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`
	IsActive    bool      `gorm:"not null" json:"is_active"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	User        User      `gorm:"foreignKey:ID"`
}

func (Track) TableName() string {
	return "track"
}
