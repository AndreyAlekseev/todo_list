package repository

import (
	todo "github.com/AndreyAlekseev/todo_list"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}
type TodoList interface {
}

type TodoItem interface {
}

type Repositiry struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repositiry {
	return &Repositiry{
		Authorization: NewAuthPostgres(db),
	}
}
