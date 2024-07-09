package domain

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	Id        string
	Title     string
	Content   string
	Status    bool
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

type TodoReq struct {
	Title string
	Content string `json:"content"`
}
