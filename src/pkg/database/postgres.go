package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/iqbalnzls/watchcommerce/src/pkg/config"
)

func NewDatabase(config *config.DatabaseConfig) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.Password, config.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(config.MaxOpenConnections)
	db.SetMaxIdleConns(config.MinIdleConnections)
	db.SetConnMaxLifetime(time.Duration(config.ConnMaxLifeTime) * time.Second)

	return db
}
