package service

import (
	todo "github.com/AndreyAlekseev/todo_list"
	"github.com/AndreyAlekseev/todo_list/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	lisrRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		lisrRepo: listRepo,
	}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.lisrRepo.GetById(userId, listId)
	if err != nil {
		//list does not exists or does not belongs to user
		return 0, err
	}
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}
