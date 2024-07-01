package models

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Title string
	Body  string
}
