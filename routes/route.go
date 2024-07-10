package routes

import (
	todo "technical-test-be-abc/app/todo/route"
	user "technical-test-be-abc/app/user/route"
	"technical-test-be-abc/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	todoGroup := e.Group("/todos", helpers.JWTMiddleware())
	userGroup := e.Group("/users")	

	todo.TodoRoute(todoGroup, db)
	user.UserRoute(userGroup, db)

}
