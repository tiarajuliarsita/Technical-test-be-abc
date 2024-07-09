package repository

import (
	"technical-test-be-abc/app/user/todo"
	"technical-test-be-abc/domain"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todo.Repository {
	return &todoRepository{
		db: db,
	}
}

func (t *todoRepository) CreateTodo(todo domain.TodoReq) error {
	err := t.db.Create(&todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) DeleteTodo(id string) error {
	todo := domain.Todo{}
	err := t.db.Delete(&todo, id).Error
	if err != nil {
		return err
	}
	return nil

}

func (t *todoRepository) GetAllTodo() ([]domain.Todo, error) {
	var todos []domain.Todo
	err := t.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *todoRepository) GetTodoById(id string) (domain.Todo, error) {
	var todo domain.Todo
	err := t.db.First(&todo, id).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (t *todoRepository) UpdateStatusTodo(id string, status bool) error {
	var todo domain.Todo
	err := t.db.First(&todo, id).Error
	if err != nil {
		return err
	}
	todo.Status = status
	err = t.db.Save(&todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) UpdateTodo(id string, todoReq domain.TodoReq) error {
	var todo domain.Todo
	err := t.db.First(&todo, id).Error
	if err != nil {
		return err
	}
	todo.Title = todoReq.Title
	todo.Content = todoReq.Content
	err = t.db.Save(&todo).Error
	if err != nil {
		return err
	}
	return nil
}
