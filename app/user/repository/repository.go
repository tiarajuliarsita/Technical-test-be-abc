package repository

import (
	"fmt"
	"technical-test-be-abc/app/user"
	"technical-test-be-abc/domains"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) FindById(userId string) (domains.User, error) {
	user := domains.User{}
	err := u.db.Where("id = ?", userId).Find(&user).Error
	if err != nil {
		fmt.Println("err : ",err)
		return domains.User{}, err
	}
	return user, nil
}

func (u *userRepository) FindByEmail(email string) (domains.User, error) {
	user := domains.User{}
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return domains.User{}, err
	}
	return user, nil
}

func (u *userRepository) Register(user domains.User) error {
	err := u.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
