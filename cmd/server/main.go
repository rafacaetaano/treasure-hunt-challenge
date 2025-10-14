package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/api"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/models"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/repository"
	"github.com/rafacaetaano/treasure-hunt-challenge/internal/user/service"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func run() error {
	// 1) Conexão com o banco (pool do database/sql + Bun)
	dsn := "postgres://postgres:admin@localhost:5432/treasuredb?sslmode=disable"
	connector := pgdriver.NewConnector(pgdriver.WithDSN(dsn))
	sqldb := sql.OpenDB(connector)

	db := bun.NewDB(sqldb, pgdialect.New())
	defer db.Close() // será executado quando run() retornar

	// 2) Boot com timeout (não travar infinito no ping/migração)
	bootCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(bootCtx); err != nil {
		return err
	}

	// 3) Criar tabela se não existir
	if _, err := db.NewCreateTable().
		Model((*models.User)(nil)).
		IfNotExists().
		Exec(bootCtx); err != nil {
		return err
	}
	log.Println("Conexão realizada e tabela ok")

	// 4) Wiring de dependências
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)

	// 5) Router
	r := gin.Default()
	r.POST("/users", api.CreateUserHandler(userSvc))
	r.GET("/users/:id", api.GetUserByIDHandler(userSvc))
	r.GET("/users", api.GetAllUsersHandler(userSvc))
	r.DELETE("/users/:id", api.DeleteUserByIDHandler(userSvc))
	r.PUT("/users/:id", api.UpdateUserByIDHandler(userSvc))
	r.POST("/userLogin", api.Login(userSvc))

	// 6) Rodar servidor (bloqueia até parar; quando sair, defer db.Close() roda)
	return r.Run(":8080")
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err) // decide encerrar aqui; defers de run() já rodaram
	}
}
