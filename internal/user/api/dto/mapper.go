package dto

import (
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/models"
)

func ToUserResponse(user *models.User) UserResponse {

	return UserResponse{
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}
}

func ToUserResponseList(users []*models.User) []UserResponse {
	out := make([]UserResponse, 0, len(users))
	for _, u := range users {
		out = append(out, ToUserResponse(u))
	}
	return out
}
