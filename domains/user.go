package domains

import (
	"technical-test-be-abc/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	Id          string `gorm:"PrimaryKey"`
	UserName    string `gorm:"not null;unique" valid:"required~your username is required"`
	Gender      string `gorm:"not null" valid:"required~your gender is required"`
	BirthOfDate string `gorm:"not null" valid:"required~your birth of date is required"`
	Email       string `gorm:"not null;unique" valid:"required~your email is required, email~invalid email format"`
	Password    string `valid:"required~your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Phone       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type UserRegister struct {
	UserName    string `json:"user_name"`
	Gender      string `json:"gender"`
	BirthOfDate string `json:"birth_of_date"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResp struct {
	Id          string `json:"id"`
	UserName    string `json:"user_name"`
	Gender      string `json:"gender"`
	BirthOfDate string `json:"birth_of_date"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	u.Id = helpers.CreateId()
	u.Password, _ = helpers.HassPass(u.Password)
	return nil
}
