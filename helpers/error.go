package helpers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func CheckError(e echo.Context, err, sub, message string, code int) (echo.Context, bool) {
	if strings.Contains(err, sub) {
		resp := ErrorResponseJson{
			Status:  false,
			Message: message,
		}
		e.JSON(code, resp)
		return e, true
	}
	return e, false
}

func CustomErr(e echo.Context, err string) error {
	e, ok := CheckError(e, err, "your title is required", "your title is required", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "your content is required", "your content is required", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "your content is required", "your content is required", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "your birt of date is required", "your birth of date is required", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "your gender is required", "your gender is required", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "your password is required", "your password is required", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "your username is required", "your username is required", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "you don`t have access", "you don`t have access", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "uni_todos_title", "your title is already exist", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "uni_users_user_name", "username is already exist", http.StatusBadRequest)
	if ok {
		return nil
	}
	
	e, ok = CheckError(e, err, "uni_users_email", "email is already exist", http.StatusBadRequest)
	if ok {
		return nil
	}
	e, ok = CheckError(e, err, "Password has to have a minimum length of 6 characters", "Password has to have a minimum length of 6 characters", http.StatusBadRequest)
	if ok {
		return nil
	}
	

	resp := ErrorResponseJson{
		Status:  false,
		Message: "internal server error",
	}
	return e.JSON(http.StatusInternalServerError, resp)
}
