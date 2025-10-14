package repository

import (
	"context"
	"database/sql"

	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/models"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

// função  construtora para criar um novo repositório de user, tem como parâmetro *bun.DB que é a conexão com banco
// retorna um ponteiro para uma struct UserRepository
// quando tivermos funções que retornam ponteiro devemos passar no return o endereço com o &
func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db: db}
}

// r *UserRepository - diz que a função CreateUser é uma função da struct UserRepository
// CreateUser - nome da função
// ctx context.Context - parâmetro que a função espera receber
// // user *models.User - parâmetro que a função espera receber no caso a struct user
func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	//ignoramos o primeiro retorno que é o resultado da query
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	//cria uma struct do tipo User
	//preenche essa nova struct com os valores com filtro pelo ID
	// retorna um ponteiro dessa struct
	user := new(models.User)
	err := r.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	//aqui só retornamos user, sem o &, porque com o new user já é um ponteiro
	return user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User

	err := r.db.NewSelect().Model(&users).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) DeleteUserByID(ctx context.Context, id int) (sql.Result, error) {
	//((*models.User)(nil)) - é para não precisar criar uma instância de User, dessa forma entende que é apenas a model User
	response, err := r.db.NewDelete().Model((*models.User)(nil)).Where("id = ?", id).Exec(ctx)
	response.RowsAffected()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *UserRepository) UpdateUserByID(ctx context.Context, id int, user *models.User) error {
	_, err := r.db.NewUpdate().Model(user).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Login(ctx context.Context, username string, password string) (*models.User, error) {
	user := new(models.User)
	err := r.db.NewSelect().Model(user).Where("username = ?", username).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
