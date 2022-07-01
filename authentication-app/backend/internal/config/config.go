package config

import (
	"context"
	"os"
	"strings"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/clog"
)

type Config struct {
	AuthDB *Database `json:"auth_db"`
}

const (
	configFilePath = "./internal/config/config.json"
)

var (
	GlobalConfig *Config
)

func GetConfig() *Config {
	return GlobalConfig
}

func InitConfig() {
	GlobalConfig = new(Config)
	initDBConfig()
	initDBs()
}

func initDBConfig() {
	DBUrl := os.Getenv("DATABASE_URL")
	if DBUrl == "" {
		clog.Fatal(
			context.Background(),
			"cannot get DATABASE_URL from env var",
		)
	}
	split := strings.Split(DBUrl, ":")
	dbUsername := split[0]
	dbPassword := split[1]
	dbHost := split[2]
	dbName := split[3]

	GlobalConfig.AuthDB = &Database{
		Username: dbUsername,
		Password: dbPassword,
		DBName: dbName,
		Host: dbHost,
	}
}