package service

import (
	"errors"
	"technical-test-be-abc/app/todo"
	"technical-test-be-abc/domains"
)

type todoService struct {
	repo todo.Repository
}


func NewTodoService(repo todo.Repository) todo.Service {
	return &todoService{
		repo: repo,
	}
}

func (t *todoService) UpdateTodoById(id string, userId string, todoReq domains.TodoReq) error {
	todo := domains.Todo{
		Title:   todoReq.Title,
		Content: todoReq.Content,
		UserId:  userId,
	}

	err := t.repo.UpdateTodoById(id, userId, todo)
	if err != nil {
		return err
	}
	return nil
}

func (t *todoService) CreateTodo(todoReq domains.TodoReq, userId string) error {
	todo := domains.Todo{
		Title:   todoReq.Title,
		Content: todoReq.Content,
		UserId:  userId,
	}
	err := t.repo.CreateTodo(todo)
	if err != nil {
		return err
	}
	return nil
}

func (t *todoService) DeleteTodo(id, userId string) error {
	err := t.repo.DeleteTodo(id, userId)
	if err != nil {
		return err
	}
	return nil
}

func (t *todoService) GetAllTodo(userId string) ([]domains.Todo, error) {
	todos, err := t.repo.GetAllTodo(userId)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *todoService) GetTodoById(id, userId string) (domains.Todo, error) {
	todo, err := t.repo.GetTodoById(id)
	if err != nil {
		return domains.Todo{}, err
	}
	if todo.UserId != userId{
		return domains.Todo{}, errors.New("you don`t have access")
	}
	return todo, nil
}

func (t *todoService) UpdateStatusTodo(id string, status bool, userId string) error {
	err := t.repo.UpdateStatusTodo(id, status, userId)
	if err != nil {
		return err
	}
	return nil
}



