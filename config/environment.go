package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
}
