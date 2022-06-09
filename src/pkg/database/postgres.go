package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/iqbalnzls/watchcommerce/src/shared/config"
)

func NewDatabase(config *config.DatabaseConfig) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.Password, config.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(config.MaxOpenConnections)
	db.SetMaxIdleConns(config.MinIdleConnections)
	db.SetConnMaxLifetime(config.ConnMaxLifeTime)

	return db
}
