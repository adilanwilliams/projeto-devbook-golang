package database

import (
	"database/sql"
	"devbook/src/config"
	"log"

	_ "github.com/lib/pq" // Databse driver
)

// Connect returns a connection with database.
func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.ConnectionStringPGDatabase)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, nil
}
