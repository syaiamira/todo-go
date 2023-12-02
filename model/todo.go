package model

import (
	"time"
)

type Todo struct {
	ID        uint      `json:"id" gorm:"autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed" gorm:"default:false"`
}
