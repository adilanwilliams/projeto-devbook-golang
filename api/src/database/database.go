package database

import (
	"database/sql"
	"devbook/src/config"
	"log"

	_ "github.com/lib/pq" // Database driver
)

// Connect initializes and verifies a connection to the PostgreSQL database.
// It returns a sql.DB instance or an error if the connection fails.
func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.ConnectionStringPGDatabase)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		defer db.Close()
		log.Fatal(err)
	}

	return db, nil
}
