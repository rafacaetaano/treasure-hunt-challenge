package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/models"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func HashPassword(password string) (string, error) {
	// Gerar o hash com custo de 10 (quanto maior o número, mais caro o hash)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "Erro ao gerar hash da senha", err
	}
	return string(hash), nil
}

func ValidatePassword(password string) error {
	//Regex para validar a senha (mínimo de 8 caracteres, 1 maiúscula, 1 número)
	re := regexp.MustCompile(`^[A-Za-z0-9]*[A-Z]+[A-Za-z0-9]*\d+[A-Za-z0-9]*$`)

	match := re.MatchString(password)
	if !match {
		return errors.New("a senha não atende as regras solicitadas")
	}
	return nil
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {

	err := user.NewValidate()
	if err != nil {
		return err
	}

	err = ValidatePassword(user.Password)
	if err != nil {
		return err
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("fail to create user on database")
	}
	return nil
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

func (s *UserService) Login(ctx context.Context, username string, password string) (*models.User, error) {
	user, err := s.repo.Login(ctx, username, password)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

//TODO: colocar loggers nas services
