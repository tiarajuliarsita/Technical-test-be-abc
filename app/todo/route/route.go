package route

import (
	"technical-test-be-abc/app/todo/handler"
	"technical-test-be-abc/app/todo/repository"
	"technical-test-be-abc/app/todo/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TodoRoute(g *echo.Group, db *gorm.DB) {
	repo := repository.NewTodoRepository(db)
	serv := service.NewTodoService(repo)
	handler := handler.NewTodoHandler(serv)

	g.POST("", handler.CreateTodo)
	g.DELETE("/:id", handler.DeleteTodoById)
	g.GET("", handler.GetAllTodos)
	g.PUT("/:id/status", handler.UpdateStatusTodo)
	g.PUT("/:id", handler.UpdateTodoById)
	g.GET("/:id", handler.GetTodoById)

}
