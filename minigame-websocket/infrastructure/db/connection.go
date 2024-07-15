package db

import (
	"fmt"
	"os"

	"github.com/oklog/ulid/v2"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	err      error
)

func Connect() {
	database, err = gorm.Open(PostgresConnection(), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{})
	database.Create(&User{Id: ulid.Make().String(), Points: 0})
}

func SqliteConnection() gorm.Dialector {
	return sqlite.Open("infrastructure/db/minigame.sqlite")
}

func PostgresConnection() gorm.Dialector {
	var (
		host     = os.Getenv("DATABASE_HOST")
		port     = os.Getenv("DATABASE_PORT")
		user     = os.Getenv("DATABASE_USER")
		dbname   = os.Getenv("DATABASE_NAME")
		sslmode  = os.Getenv("DATABASE_SSLMODE")
		password = os.Getenv("DATABASE_PASSWORD")
	)

	var connectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	return postgres.Open(connectionString)
}
