package database

import "stepback-golang/internal/model"

func Migrate() {
	DB.AutoMigrate(
		&model.User{},
		&model.Address{},
	)
}
