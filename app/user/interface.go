package user

import (
	"technical-test-be-abc/domains"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		Register(user domains.User) error
		FindByEmail(email string) (domains.User, error)
		FindById(userId string)(domains.User, error)
		
	}

	Service interface {
		Register(user domains.UserRegister) error
		Login(userReq domains.UserLogin) (string, error)
		GetProfile(userId string)(domains.User, error)
	}

	Handler interface {
		Register(e echo.Context) error
		Login(e echo.Context) error
		GetProfile(e echo.Context)error
	}
)
