package main

import (
	"log"
	"net/http"

	"stepback-golang/internal/config"
	"stepback-golang/internal/database"
	"stepback-golang/internal/repository"
	"stepback-golang/internal/router"
	"stepback-golang/internal/service"
	"stepback-golang/internal/utils"
)

func main() {
	cfg := config.Load()

	database.Connect(cfg.DB)
	database.Migrate()
	defer database.Close()

	utils.InitJWT(&cfg.JWT)

	userRepo := repository.NewUserRepository(database.DB)
	authService := service.NewAuthService(userRepo)
	r := router.NewWithDeps(authService)

	addr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Println("Server starting on " + addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
