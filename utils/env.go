package utils

import (
  "fmt"
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadENV() error {
	token := os.Getenv("TURSO_AUTH_TOKEN")

	if token == "" {
		fmt.Println("trying to load .env")
		err := godotenv.Load(".env")

		if err != nil {
			return errors.New("failed to load environment variables from .env")
		}
	}

	return nil
}
