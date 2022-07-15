package database

import (
	"database/sql"
	"fmt"
	"stream-it-consulting/innovation-team/example-rest-api/config"
	"time"

	_ "github.com/lib/pq"
)

func Initialize() (db *sql.DB) {
	db, err := sql.Open("postgres", config.AppConfig.DatabaseDSN)

	if err != nil {
		fmt.Printf("Cannot connect to database")
	}

	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxOpenConns(200)

	return db
}
