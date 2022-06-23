package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/clog"
)

type Config struct {
	AuthDB *Database `json:"auth_db"`
}

const (
	configFilePath = "../config/config.json"
)

var (
	GlobalConfig *Config
	once         sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		initConfig()
	})

	return GlobalConfig
}

func initConfig() {
	configFile, _ := os.Open(configFilePath)
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)

	var configuration Config
	err := decoder.Decode(&configuration)
	if err != nil {
		clog.Fatal(
			context.Background(),
			fmt.Sprintf("failed to load configs err=% s", err.Error()),
		)
	}

	GlobalConfig = &configuration
}
