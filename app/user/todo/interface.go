package todo

import "technical-test-be-abc/domain"

type Repository interface {
	CreateTodo(todo domain.TodoReq) error
	GetAllTodo() ([]domain.Todo, error)
	GetTodoById(id string) (domain.Todo, error)
	DeleteTodo(id string) error
	UpdateTodo(id string, todo domain.TodoReq) error
	UpdateStatusTodo(id string, status bool) error
}
type Service interface {
	CreateTodo(todo domain.TodoReq) error
	GetAllTodo() ([]domain.Todo, error)
	GetTodoById(id string) (domain.Todo, error)
	DeleteTodo(id string) error
	UpdateTodo(id string, todo domain.TodoReq) error
	UpdateStatusTodo(id string, status bool) error
}
