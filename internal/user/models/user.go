package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	//define o nome da tabela
	bun.BaseModel `bun:"table:users"`

	ID        uint      `bun:"id,pk,autoincrement"` //primary_key e autoincremento na tabela
	Name      string    `bun:"name,notnull"`
	Username  string    `bun:"username,notnull,unique"`
	Password  string    `bun:"password,notnull"`
	Email     string    `bun:"email,notnull,unique"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"` //se não for passado nenhum valor o Postgres preenche automaticamente
	Role      string    `bun:"role,notnull"`
}
