package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Provider string `json:"provider"`
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Todos    []Todo `gorm:"foreignKey:UserID;references:ID"`
}