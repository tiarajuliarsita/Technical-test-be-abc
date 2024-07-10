package repository

import (
	"errors"
	"fmt"
	"technical-test-be-abc/app/todo"
	"technical-test-be-abc/domains"

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

func (t *todoRepository) CreateTodo(todo domains.Todo) error {
	err := t.db.Create(&todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) DeleteTodo(id, userId string) error {
	todo := domains.Todo{}
	todo, err := t.GetTodoById(id)
	if err != nil {
		fmt.Println("err :  ", err)
		return err
	}

	if todo.UserId != userId {
		return errors.New("you don`t have access")
	}

	err = t.db.Where("id = ?", id).Delete(&todo).Error
	if err != nil {
		return err
	}
	return nil

}

func (t *todoRepository) GetAllTodo(userId string) ([]domains.Todo, error) {
	var todos []domains.Todo
	err := t.db.Where("user_id = ?", userId).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *todoRepository) GetTodoById(id string) (domains.Todo, error) {
	var todo domains.Todo
	err := t.db.Where("id = ?", id).Find(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (t *todoRepository) UpdateStatusTodo(id string, status bool, userId string) error {
	var todo domains.Todo
	todo, err := t.GetTodoById(id)
	if err != nil {
		return err
	}
	if todo.UserId != userId {
		return errors.New("you don`t have access")
	}
	todo.Status = status
	err = t.db.Save(&todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) UpdateTodoById(id, userId string, todo domains.Todo) error {
	result, err := t.GetTodoById(id)
	if err != nil {
		return err
	}

	if result.UserId != userId {
		return errors.New("you don`t have access to delete this data")
	}

	err = t.db.Where("id =?", id).Updates(&todo).Error
	if err != nil {
		return err
	}
	return nil
}
