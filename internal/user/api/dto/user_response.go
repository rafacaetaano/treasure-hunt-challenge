package dto

import "time"

type UserResponse struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Role      string    `json:"role"`
}
