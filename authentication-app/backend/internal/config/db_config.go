package config

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)

type Database struct {
	Username       string `json:"db_username"`
	Password       string `json:"db_password"`
	Port           string `json:"db_port"`
	DBName         string `json:"db_name"`
	PoolMaxOpen    int    `json:"pool_max_open"`
	PoolMaxIdle    int    `json:"pool_max_idle"`
	MaxIdleSeconds int    `json:"max_idle_seconds"`
	MaxLifeSeconds int    `json:"max_life_seconds"`
}

type dbs struct {
	AuthDB *sql.DB
}

var (
	globalDBs *dbs
	dbOnce    sync.Once
)

func GetDBs() *dbs {
	dbOnce.Do(func() {
		initDBs()
	})

	return globalDBs
}

func initDBs() {
	initAuthDB()
}

func initAuthDB() {
	pool, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(localhost:%s)/%s?parseTime=true",
			GetConfig().AuthDB.Username,
			GetConfig().AuthDB.Password,
			GetConfig().AuthDB.Port,
			GetConfig().AuthDB.DBName))
	if err != nil {
		log.Fatal("start db error", err.Error())
	}

	if err = pool.Ping(); err != nil {
		log.Fatal("reach db error", err.Error())
	}

	pool.SetMaxIdleConns(GetConfig().AuthDB.PoolMaxIdle)
	pool.SetMaxOpenConns(GetConfig().AuthDB.PoolMaxOpen)
	pool.SetConnMaxIdleTime(
		time.Duration(GetConfig().AuthDB.MaxIdleSeconds) * time.Second)
	pool.SetConnMaxLifetime(
		time.Duration(GetConfig().AuthDB.MaxLifeSeconds) * time.Second)

	globalDBs.AuthDB = pool
}
