package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repositiry{}
}
