package repository

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

func NewRepository() *Repositiry {
	return &Repositiry{}
}
