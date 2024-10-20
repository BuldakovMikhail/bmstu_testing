package main

import (
	"github.com/rs/zerolog/log"
	postgres2 "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"src/cmd/muzyaka"
)

// @title Muzyaka API
// @version 1.0
// @description API Server for musical service

// @host localhost:8080
// @BasePath /

// @in header
// @name Authorization

func main() {
	dsn := "host=localhost user=postgres password=123 dbname=postgres port=5432"
	db, err := gorm.Open(postgres2.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error().Err(err)
	}
	if err != nil {
		log.Error().Err(err)
	}

	muzyaka.App(db)
}
