package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


type Config struct {
	AppPort string

	PostgresHost string
	PostgresUser string
	PostgresPassword string
	PostgresPort string
	PostgresDB string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	JWTSecret   string
	JWTExpireIn int

}

func LoadConfig() *Config{
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("[WARN] .env file not found, using system environment variables")
	}
	// redisDB , err := strconv.Atoi(GetRequrement("Redis_DB"))
	// if err != nil {
	// 	log.Fatal("Redis_DB is not found in .env")
	// }
	JWTExpireIn, err := strconv.Atoi(GetRequrement("JWTExpireIn"))
	if err != nil {
		log.Fatal("JWTExpireIn is not found in .env")
	}
	return &Config{
		AppPort: GetRequrement("AppPort"),
		PostgresHost: GetRequrement("PostgresHost"),
		PostgresUser: GetRequrement("PostgresUser"),
		PostgresPassword: GetRequrement("PostgresPassword"),
		PostgresPort: GetRequrement("PostgresPort"),
		PostgresDB: GetRequrement("PostgresDB"),

		JWTSecret: GetRequrement("JWTSecret"),
		JWTExpireIn: JWTExpireIn,
	}

}

func GetRequrement(key string) string {
	if value , exist := os.LookupEnv(key); exist {
		return value
	}
	log.Fatal("Required environment variable %s is not set", key)
	return ""
}
