package util

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	envFilePath = ".env"
)

func GetDotEnvVariable(key string) (string, error) {
	err := godotenv.Load(envFilePath)
	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}
