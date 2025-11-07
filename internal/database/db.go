package database

import (
	"fmt"
	"log"
	"recsys/config"
	"recsys/internal/movie"
	"recsys/internal/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB * gorm.DB


func InitDatabase(cfg *config.Config) *gorm.DB{
	dsn := fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
		cfg.PostgresPort,
	)
	var err error

	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to Postgres: %v", err)
	}
	log.Println("✅ Postgres connected")
	err = db.AutoMigrate(&user.User{}, &movie.Movie{})
	if err != nil {
		log.Fatal("migrate is failed")
	}
	log.Println("Database is migrated")

	db = DB
	return db
}
