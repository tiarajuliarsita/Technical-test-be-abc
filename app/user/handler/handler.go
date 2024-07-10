package handler

import (
	"fmt"
	"net/http"
	"technical-test-be-abc/app/user"
	"technical-test-be-abc/domains"
	"technical-test-be-abc/helpers"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	serv user.Service
}

func NewUserHandler(serv user.Service) user.Handler {
	return &userHandler{serv: serv}
}

func (u *userHandler) GetProfile(e echo.Context) error {
	userId, _, err := helpers.ExtractToken(e)
	if err != nil {
		fmt.Println("err : ",err)

		return helpers.CustomErr(e, err.Error())
	}
	user, err := u.serv.GetProfile(userId)
	if err != nil {
		fmt.Println("err : ",err)

		return helpers.CustomErr(e, err.Error())
	}
	userResp := domains.UserToUserResp(user)
	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully got the profile", userResp))

}

func (u *userHandler) Register(e echo.Context) error {
	userReq := domains.UserRegister{}
	err := e.Bind(&userReq)
	if err != nil {
		fmt.Println("err", err)
		return helpers.CustomErr(e, err.Error())
	}

	err = u.serv.Register(userReq)
	if err != nil {
		fmt.Println("err", err)

		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusCreated, helpers.SuccessResponse("registered successfully", nil))
}

func (u *userHandler) Login(e echo.Context) error {
	userReq := domains.UserLogin{}
	err := e.Bind(&userReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	token, err := u.serv.Login(userReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully login", echo.Map{
		"token": token,
	}))
}
