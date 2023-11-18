package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
}
