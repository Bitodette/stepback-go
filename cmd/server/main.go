package main

import (
	"log"

	"stepback-golang/internal/config"
	"stepback-golang/internal/database"
)

func main() {
	cfg := config.Load()

	database.Connect(cfg.DB)
	database.Migrate()
	defer database.Close()

	log.Println("Server starting on " + cfg.Server.Host + ":" + cfg.Server.Port)
}
