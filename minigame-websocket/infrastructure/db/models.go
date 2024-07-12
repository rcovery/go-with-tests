package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id     string
	Points int
}
