package models

import (
	"fmt"
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

func (u *User) NewValidate() error {
	if u.Email == "" {
		return fmt.Errorf("o email deve ser informado")
	}
	if u.Name == "" {
		return fmt.Errorf("o nome deve ser informado")
	}
	if u.Username == "" {
		return fmt.Errorf("o usuário deve ser informado")
	}
	if u.Password == "" {
		return fmt.Errorf("a senha deve ser informada")
	}
	if u.Role == "" {
		return fmt.Errorf("a função deve ser informada")
	}
	return nil
}
