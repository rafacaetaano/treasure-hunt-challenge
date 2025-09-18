package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/api"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/models"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/repository"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/service"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"golang.org/x/net/context"
)

func main() {
	// -------------------------------
	// Conexão com banco de dados
	// -------------------------------
	dsn := "postgres://postgres:admin@localhost:5432/treasuredb?sslmode=disable"
	connector := pgdriver.NewConnector(pgdriver.WithDSN(dsn))
	sqldb := sql.OpenDB(connector)
	db := bun.NewDB(sqldb, pgdialect.New())

	defer func() {
		if err := db.Close(); err != nil {
			log.Println("Erro ao fechar a conexão:", err)
		}
	}()

	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao conectar no banco", err)
	}
	fmt.Println("Conexão com banco realizada com sucesso")

	//Cria tabela

	ctx := context.Background()

	if _, err := db.NewCreateTable().
		Model((*models.User)(nil)).
		IfNotExists().
		Exec(ctx); err != nil {
		log.Fatal("Erro ao criar tabela users:", err)
	}

	// Repository, Service
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)

	// Router
	r := gin.Default()
	r.POST("/users", api.CreateUserHandler(userSvc))
	r.GET("users/:id", api.GetUserByIDHandler(userSvc))
	r.GET("/users", api.GetAllUsers(userSvc))
	r.DELETE("/users/:id", api.DeleteUserByIDHandler(userSvc))

	// Rodar servidor
	r.Run(":8080")
}
