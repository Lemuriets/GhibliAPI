package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey" json:"id"`
	Login    string `gorm:"unique; not null" json:"login"`
	Password string `gorm:"not null" json:"password"`
	Nickname string `json:"nickname"`
}

type Film struct {
	gorm.Model

	ID uint `gorm:"primarykey" json:"id"`
}
