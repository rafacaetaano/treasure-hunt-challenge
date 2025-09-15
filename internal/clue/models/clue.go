package models

import (
	"time"

	track "github.com/rafacaetaano/treasure-hunt-challenge/internal/track/models"
)

type Clue struct {
	ID           uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	ClueText     string      `gorm:"not null" json:"clue_text"`
	Track        track.Track `gorm:"foreignKey:ID"`
	Date         time.Time   `gorm:"not null" json:"clue_date"`
	NextClueDate time.Time   `gorm:"not null" json:"next_clue_date"`
	IsActive     bool        `gorm:"not null" json:"is_active"`
	Position     int         `gorm:"not null" json:"position"`
}

func (Clue) TableName() string {
	return "clue"
}
