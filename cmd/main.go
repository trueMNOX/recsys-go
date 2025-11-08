package main

import (
	"log"
	"recsys/config"
	"recsys/internal/database"
	"recsys/internal/movie"
	"recsys/internal/user"

	"github.com/gin-gonic/gin"
)

func main(){

	cfg := config.LoadConfig()
	Database := database.InitDatabase(cfg)

	userRepo := user.NewUserRepository(Database)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHanler(userService)

	movieRepo := movie.NewMovieRepository(Database)
	movieService := movie.NewMovieService(movieRepo, userRepo)
	movieHandler := movie.NewMovieHandler(movieService)


	r := gin.Default()
	log.Println("Starting recsys server...")

	v1 := r.Group("/api/v1")
	movieHandler.MovieRoute(v1)
	v2 := r.Group("/api/v2")
	userHandler.UserRoute(v2)

	r.Run(":8080")
}
