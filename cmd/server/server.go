package server

import (
	"gorm.io/gorm"
	"log"
	"zuri-stage-one/internal/adapters/repository"
	"zuri-stage-one/internal/core/helpers"
	"zuri-stage-one/internal/core/models"
)

func Run() (*gorm.DB, error) {
	err := helpers.Load()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db, err := repository.ConnectDb(&helpers.Config{
		DBUser:     helpers.Instance.DBUser,
		DBPass:     helpers.Instance.DBPass,
		DBHost:     helpers.Instance.DBHost,
		DBName:     helpers.Instance.DBName,
		DBPort:     helpers.Instance.DBPort,
		DBTimeZone: helpers.Instance.DBTimeZone,
		DBMode:     helpers.Instance.DBMode,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = repository.MigrateAll(db)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	user := &models.User{
		SlackUsername: "tchisom17",
		Backend:       true,
		Age:           40,
		Bio:           "Just the guy!",
	}

	db.Create(user)
	return db, nil
}
