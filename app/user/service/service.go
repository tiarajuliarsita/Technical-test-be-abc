package service

import (
	"errors"
	"technical-test-be-abc/app/user"
	"technical-test-be-abc/domains"
	"technical-test-be-abc/helpers"
)

type userService struct {
	repo user.Repository
}

func (u *userService) GetProfile(userId string) (domains.User, error) {
	user, err := u.repo.FindById(userId)
	if err != nil {
		return domains.User{}, err
	}
	return user, nil
}

func NewUserService(repo user.Repository) user.Service {
	return &userService{
		repo: repo,
	}
}

func (u *userService) Register(userReq domains.UserRegister) error {
	user := domains.User{
		UserName:    userReq.UserName,
		Gender:      userReq.Gender,
		Password:    userReq.Password,
		Phone:       userReq.Phone,
		Email:       userReq.Email,
		BirthOfDate: userReq.BirthOfDate,
	}

	err := u.repo.Register(user)
	if err != nil {
		return err
	}
	return nil
}

// Login implements user.Service.
func (u *userService) Login(userReq domains.UserLogin) (string, error) {
	user, err := u.repo.FindByEmail(userReq.Email)
	if err != nil {
		return "", err
	}
	ok, err := helpers.ComparePass([]byte(user.Password), []byte(userReq.Password))
	if err != nil {
		return "", err
	}

	if !ok {
		if err != nil {
			return "", errors.New("invalid password")
		}
	}

	token, err := helpers.CreateToken(user.Id, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil

}
