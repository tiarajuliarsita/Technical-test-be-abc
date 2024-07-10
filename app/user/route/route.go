package route

import (
	"technical-test-be-abc/app/user/handler"
	"technical-test-be-abc/app/user/repository"
	"technical-test-be-abc/app/user/service"
	"technical-test-be-abc/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoute(g *echo.Group, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	serv := service.NewUserService(repo)
	handler := handler.NewUserHandler(serv)

	g.POST("/register", handler.Register)
	g.POST("/login", handler.Login)

	auth := g.Group("/profile", helpers.JWTMiddleware())
	auth.GET("", handler.GetProfile)
}

