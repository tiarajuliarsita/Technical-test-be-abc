package handler

import (
	"fmt"
	"net/http"
	"technical-test-be-abc/app/todo"
	"technical-test-be-abc/domains"
	"technical-test-be-abc/helpers"

	"github.com/labstack/echo/v4"
)

type todoHandler struct {
	serv todo.Service
}

func NewTodoHandler(serv todo.Service) todo.Handler {
	return &todoHandler{serv: serv}
}

func (t *todoHandler) CreateTodo(e echo.Context) error {
	id, _, err := helpers.ExtractToken(e)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}
	todoReq := domains.TodoReq{}
	err = e.Bind(&todoReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	err = t.serv.CreateTodo(todoReq, id)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusCreated, helpers.SuccessResponse("success create todo", nil))
}

func (t *todoHandler) DeleteTodoById(e echo.Context) error {

	todoId := e.Param("id")
	id, _, err := helpers.ExtractToken(e)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	err = t.serv.DeleteTodo(todoId, id)
	if err != nil {
		fmt.Println("err", err)
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully deleted",nil))
}

func (t *todoHandler) GetAllTodos(e echo.Context) error {
	id, _, err := helpers.ExtractToken(e)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	todos, err := t.serv.GetAllTodo(id)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	todoResp := domains.RespTodos(todos)

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully got the data", todoResp))
}

func (t *todoHandler) UpdateStatusTodo(e echo.Context) error {
	status := domains.Status{}
	todoId := e.Param("id")

	userId, _, err := helpers.ExtractToken(e)
	if err != nil {
		fmt.Println("err h", err)

		return helpers.CustomErr(e, err.Error())
	}

	err = e.Bind(&status)
	if err != nil {
		fmt.Println("err h", err)

		return helpers.CustomErr(e, err.Error())
	}

	err = t.serv.UpdateStatusTodo(todoId, status.Status, userId)
	if err != nil {
		fmt.Println("err h", err)
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully updated status", nil))

}

func (t *todoHandler) UpdateTodoById(e echo.Context) error {
	todoReq := domains.TodoReq{}
	todoId := e.Param("id")

	userId, _, err := helpers.ExtractToken(e)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	err = e.Bind(&todoReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	err = t.serv.UpdateTodoById(todoId, userId, todoReq)
	if err != nil {
		
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully updated todo", nil))

}

func (t *todoHandler) GetTodoById(e echo.Context) error {
	todoId := e.Param("id")

	userId, _, err := helpers.ExtractToken(e)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	todo, err := t.serv.GetTodoById(todoId, userId)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}
	todoResp := domains.TodoToTodoResp(todo)

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully got the todo", todoResp))

}
