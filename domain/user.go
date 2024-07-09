package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id          string
	UserName    string
	Gender      string
	BirthOfDate string
	Email       string
	Password    string
	Phone       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeleteAt    gorm.DeletedAt `gorm:"index"`
}

type UserRegister struct {
    UserName    string `json:"user_name"`
    Gender      string `json:"gender"`
    BirthOfDate string `json:"birth_of_date"`
    Email       string `json:"email"`
    Password    string `json:"password"`
    Phone       string `json:"phone"`
}
