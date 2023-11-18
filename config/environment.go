package config

import (

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
}
