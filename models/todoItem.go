package models

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Title string `gorm:"not null"`
	Body  string `gorm:"not null"`
}
