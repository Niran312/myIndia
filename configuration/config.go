package configuration

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Env          string    `env:"ENV"`
	Mysql        MysqlData `json:"mysql"`
	Host         string    `env:"APP_HOST"`
	Port         string    `env:"APP_PORT"`
	JWTSecretKey string    `env:"JWT_ACCESS_SIGN_KEY"`
}

func LoadConfig() {
	currentPath, err := os.Getwd()

	log.Printf("Current Path: %v", currentPath)

	if err != nil {
		log.Println(err)
	}

	environmentPath := filepath.Join(currentPath, ".env")

	log.Printf("Environment Path: %v", environmentPath)

	if err := godotenv.Load(environmentPath); err != nil {
		log.Fatal("Error loading .env file \n", err)
	}
}

func GetConfig() Config {
	return Config{
		Env:          os.Getenv("ENV"),
		Mysql:        GetMsqlData(),
		Host:         os.Getenv("APP_HOST"),
		Port:         os.Getenv("APP_PORT"),
		JWTSecretKey: os.Getenv("JWT_ACCESS_SIGN_KEY"),
	}
}
