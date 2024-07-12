package db

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {
	database, _ := gorm.Open(sqlite.Open("infrastructure/db/minigame.sqlite"), &gorm.Config{})
	database.AutoMigrate(&User{})
	database.Create(&User{Id: ulid.Make().String(), Points: 0})
}
