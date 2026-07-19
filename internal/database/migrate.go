package database

import "stepback-golang/internal/model"

// auto creates tables from models
// ok for dev, for prod pake sql migration files
func Migrate() {
	DB.AutoMigrate(
		&model.User{},
		&model.Address{},
	)
}
