package service

import (
	"context"
	"errors"
	"log"

	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/models"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return s.repo.GetAllUsers(ctx)
}

func (s *UserService) DeleteUserByID(ctx context.Context, id int) error {
	response, err := s.repo.DeleteUserByID(ctx, id)
	if err != nil {
		log.Println("Error deleting user")
		return err
	}

	result, err := response.RowsAffected()
	if err != nil {
		log.Println("Error get rows affected")
		return err
	}

	if result == 0 {
		log.Println("Nenhum usuário encontrado")
		return errors.New("nenhum usuário encontrado")
	}

	return nil

}

func (s *UserService) UpdateUserByID(ctx context.Context, id int, user *models.User) error {
	return s.repo.UpdateUserByID(ctx, id, user)
}

//TODO: colocar loggers nas services
