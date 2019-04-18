package db

import (
	"database/sql"
	"fmt"
)

// Postgres represents a Postgres database by wrapping an SQL database handle.
type Postgres struct {
	db
	DB *sql.DB
}

// Open opens a connection to a Postgres database based on the configuration,
// and pings it to check the connection.
//
// The host, port, user, and name are required. If the user has a password, it
// is also required.
func (db *Postgres) Open() error {
	dsName := db.buildDataSourceName()

	handle, err := sql.Open("postgres", dsName)
	if err != nil {
		return fmt.Errorf("db.Postgres.Open: failed to open database (%s)", err)
	}

	db.DB = handle

	if err := db.DB.Ping(); err != nil {
		return fmt.Errorf("db.Postgres.Open: failed to ping database (%s)", err)
	}

	return nil
}

// Close closes the connection to the Postgres database.
func (db *Postgres) Close() error {
	if db.DB == nil {
		return nil
	}

	if err := db.DB.Close(); err != nil {
		return fmt.Errorf("db.Postgres.Close: failed to close database (%s)", err)
	}

	db.DB = nil

	return nil
}

func (db *Postgres) buildDataSourceName() string {
	dsName := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s",
		db.conf.Host,
		db.conf.Port,
		db.conf.Name,
		db.conf.User,
	)

	if db.conf.Password != "" {
		dsName += fmt.Sprintf(" password=%s", db.conf.Password)
	}

	return dsName
}
