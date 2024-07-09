package service

import (
	"technical-test-be-abc/app/user/todo"
	"technical-test-be-abc/domain"
)

type todoService struct {
	repo todo.Repository
}

func NewService(todo todo.Repository) todo.Service {
	return &todoService{
		repo: todo,
	}
}

func (t *todoService) CreateTodo(todo domain.TodoReq) error {
	err := t.repo.CreateTodo(todo)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTodo implements todo.Service.
func (t *todoService) DeleteTodo(id string) error {
	err := t.repo.DeleteTodo(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAllTodo implements todo.Service.
func (t *todoService) GetAllTodo() ([]domain.Todo, error) {
	todos, err := t.repo.GetAllTodo()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// GetTodoById implements todo.Service.
func (t *todoService) GetTodoById(id string) (domain.Todo, error) {
	todo, err := t.repo.GetTodoById(id)
	if err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func (t *todoService) UpdateStatusTodo(id string, status bool) error {
	err := t.repo.UpdateStatusTodo(id, status)
	if err != nil {
		return err
	}
	return nil
}

func (t *todoService) UpdateTodo(id string, todo domain.TodoReq) error {
	err := t.repo.UpdateTodo(id, todo)
	if err != nil {
		return err
	}
	return nil
}
