package todo

import (
	"technical-test-be-abc/domains"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		CreateTodo(todo domains.Todo) error
		GetAllTodo(userId string) ([]domains.Todo, error)
		GetTodoById(id string) (domains.Todo, error)
		DeleteTodo(id, userId string) error
		UpdateTodoById(id, userId string, todo domains.Todo) error
		UpdateStatusTodo(id string, status bool, userId string) error

	}

	Service interface {
		CreateTodo(todo domains.TodoReq, userId string) error
		GetAllTodo(UserId string) ([]domains.Todo, error)
		GetTodoById(id, userId string) (domains.Todo, error)
		DeleteTodo(id, userId string) error
		UpdateStatusTodo(id string, status bool, userId string) error
		UpdateTodoById(id, userId string, todoReq domains.TodoReq) error

	}

	Handler interface {
		CreateTodo(e echo.Context) error
		DeleteTodoById(e echo.Context) error
		GetAllTodos(e echo.Context) error
		UpdateStatusTodo(e echo.Context) error
		UpdateTodoById(e echo.Context) error
		GetTodoById(e echo.Context) error
	}
)
