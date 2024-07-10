package service

import (
	"errors"
	"technical-test-be-abc/app/todo/mocks"
	"technical-test-be-abc/domains"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateTodoById(t *testing.T) {
	mockRepo := new(mocks.Repository)
	todoService := NewTodoService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		todoReq := domains.TodoReq{
			Title:   "Updated Title",
			Content: "Updated Content",
		}
		mockRepo.On("UpdateTodoById", "todo-1", "user-1", mock.Anything).Return(nil).Once()

		err := todoService.UpdateTodoById("todo-1", "user-1", todoReq)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		todoReq := domains.TodoReq{
			Title:   "Updated Title",
			Content: "Updated Content",
		}
		mockRepo.On("UpdateTodoById", "todo-1", "user-1", mock.Anything).Return(errors.New("repository error")).Once()

		err := todoService.UpdateTodoById("todo-1", "user-1", todoReq)
		assert.Error(t, err)
		assert.Equal(t, "repository error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func CreateTodo(t *testing.T) {
	mockRepo := new(mocks.Repository)
	todoService := NewTodoService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		todoReq := domains.TodoReq{
			Title:   "Test Title",
			Content: "Test Content",
		}
		userId := "user-1"

		mockRepo.On("CreateTodo", mock.AnythingOfType("domains.Todo")).Return(nil).Once()

		err := todoService.CreateTodo(todoReq, userId)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		todoReq := domains.TodoReq{
			Title:   "Test Title",
			Content: "Test Content",
		}
		userId := "user-1"

		mockRepo.On("CreateTodo", mock.AnythingOfType("domains.Todo")).Return(errors.New("repository error")).Once()

		err := todoService.CreateTodo(todoReq, userId)
		assert.Error(t, err)
		assert.Equal(t, "repository error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteTodo(t *testing.T) {
	mockRepo := new(mocks.Repository)
	todoService := NewTodoService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("DeleteTodo", "123", "user-1").Return(nil).Once()

		err := todoService.DeleteTodo("123", "user-1")
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.On("DeleteTodo", "123", "user-1").Return(errors.New("repository error")).Once()

		err := todoService.DeleteTodo("123", "user-1")
		assert.Error(t, err)
		assert.Equal(t, "repository error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetAllTodo(t *testing.T) {
	mockRepo := new(mocks.Repository)
	todoService := NewTodoService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		expectedTodos := []domains.Todo{
			{Id: "1", Title: "Todo 1", Content: "Content 1", UserId: "user-1"},
			{Id: "2", Title: "Todo 2", Content: "Content 2", UserId: "user-1"},
		}
		mockRepo.On("GetAllTodo", "user-1").Return(expectedTodos, nil).Once()

		todos, err := todoService.GetAllTodo("user-1")
		assert.NoError(t, err)
		assert.Equal(t, expectedTodos, todos)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.On("GetAllTodo", "user-1").Return(nil, errors.New("repository error")).Once()

		todos, err := todoService.GetAllTodo("user-1")
		assert.Error(t, err)
		assert.Nil(t, todos)
		assert.Equal(t, "repository error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetTodoById(t *testing.T) {
	mockRepo := new(mocks.Repository)
	todoService := NewTodoService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		expectedTodo := domains.Todo{
			Id:     "1",
			Title:  "Todo 1",
			UserId: "user-1",
		}
		mockRepo.On("GetTodoById", "1").Return(expectedTodo, nil).Once()

		todo, err := todoService.GetTodoById("1", "user-1")
		assert.NoError(t, err)
		assert.Equal(t, expectedTodo, todo)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Unauthorized Access", func(t *testing.T) {
		unauthorizedTodo := domains.Todo{
			Id:     "2",
			Title:  "Todo 2",
			UserId: "user-2", 
		}
		mockRepo.On("GetTodoById", "2").Return(unauthorizedTodo, nil).Once()

		todo, err := todoService.GetTodoById("2", "user-1")
		assert.Error(t, err)
		assert.Equal(t, "you don`t have access", err.Error())
		assert.Empty(t, todo) 
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.On("GetTodoById", "3").Return(domains.Todo{}, errors.New("repository error")).Once()

		todo, err := todoService.GetTodoById("3", "user-1")
		assert.Error(t, err)
		assert.Equal(t, "repository error", err.Error())
		assert.Empty(t, todo) 
		mockRepo.AssertExpectations(t)
	})
}


func TestTodoService_UpdateStatusTodo(t *testing.T) {
	mockRepo := new(mocks.Repository)
	todoService := NewTodoService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("UpdateStatusTodo", "1", true, "user-1").Return(nil).Once()

		err := todoService.UpdateStatusTodo("1", true, "user-1")
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Unauthorized Access", func(t *testing.T) {
		mockRepo.On("UpdateStatusTodo", "2", true, "user-2").Return(errors.New("you don't have access")).Once()

		err := todoService.UpdateStatusTodo("2", true, "user-2")
		assert.Error(t, err)
		assert.Equal(t, "you don't have access", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.On("UpdateStatusTodo", "3", true, "user-1").Return(errors.New("repository error")).Once()

		err := todoService.UpdateStatusTodo("3", true, "user-1")
		assert.Error(t, err)
		assert.Equal(t, "repository error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

