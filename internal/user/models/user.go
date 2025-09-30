package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	//define o nome da tabela
	bun.BaseModel `bun:"table:users"`

	ID        uint      `bun:"id,pk,autoincrement" json:"id"` //primary_key e autoincremento na tabela
	Name      string    `bun:"name,notnull" json:"name"`
	Username  string    `bun:"username,notnull,unique" json:"username"`
	Password  string    `bun:"password,notnull" json:"password"`
	Email     string    `bun:"email,notnull,unique" json:"email"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp" json:"created_at"` //se não for passado nenhum valor o Postgres preenche automaticamente
	Role      string    `bun:"role,notnull" json:"role"`
}
