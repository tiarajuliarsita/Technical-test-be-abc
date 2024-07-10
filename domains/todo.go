package domains

import (
	"technical-test-be-abc/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Todo struct {
	Id        string
	Title     string `gorm:"not null;unique" valid:"required~your title is required"`
	Content   string `gorm:"not null" valid:"required~your content is required"`
	Status    bool   `gorm:"default:true"`
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type TodoReq struct {
	Title   string
	Content string `json:"content"`
}

type Status struct {
	Status bool `json:"status"`
}

type TodoResp struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    bool      `json:"status"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Todo) BeforeCreate(tx *gorm.DB) (error) {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}
	t.Id = helpers.CreateId()
	return nil
}
