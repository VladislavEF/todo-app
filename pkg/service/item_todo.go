package service

import (
	"github.com/VladislavEF/todo-app"
	"github.com/VladislavEF/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listrepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listrepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listrepo: listrepo}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listrepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}
