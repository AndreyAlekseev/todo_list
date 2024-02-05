package service

import "github.com/AndreyAlekseev/todo_list/pkg/repository"

type Authorization interface {
}
type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repositiry) *Service {
	return &Service{}
}
