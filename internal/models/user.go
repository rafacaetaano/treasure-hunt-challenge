package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Username  string    `gorm:"not null;unique" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	Email     string    `gorm:"not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	Role      string    `gorm:"not null" json:"role"`
}

// TableName define explicitamente o nome da tabela no banco
func (User) TableName() string {
	return "users"
}
