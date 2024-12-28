package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetPort() string {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		return "8000"
	}

	return port
}

func InitialEnv(path string) {
	if err := godotenv.Load(path); err != nil {
		log.Println("No .env file found")
	}
}
