package db

import (
	"fmt"
)

const (
	// TypeInMemory represents a database that persists entities in memory.
	TypeInMemory = "inmemory"

	// TypePostgres represents a Postgres database.
	TypePostgres = "postgres"
)

// Config represents the information required to establish a database
// connection.
//
// Not all fields are required, depending on the type of database. For example,
// an in memory database does not require anything.
type Config struct {
	// Host represents the address at which the database is hosted.
	//
	// For example, if the database is hosted on the local machine, it would be
	// "localhost".
	Host string

	// Port represents the port to use when connecting to the database on the
	// given host.
	Port string

	// User represents the name of the database user to use to connect to the
	// database.
	User string

	// Password represents database user's password.
	//
	// If no password is required, it can be left empty.
	Password string

	// Name represents the name of the database to use.
	Name string

	// SSLMode represents whether or not SSL is required to connect to the
	// database.
	SSLMode string
}

// A DB is an interface representing the ability to open and close a connection
// to a database.
type DB interface {
	// Open opens a database connection and checks it with a ping.
	Open() error

	// Close closes a database connection.
	Close() error
}

type db struct {
	conf Config
}

// New creates a database of the given type based on the configuration.
//
// If the database type is not recognized, it returns an error.
//
// The database must be opened before it can be used, and ideally closed when
// it is no longer needed.
func New(dbType string, dbConf Config) (DB, error) {
	switch dbType {
	case TypeInMemory:
		return NewInMemory(dbConf), nil
	case TypePostgres:
		return NewPostgres(dbConf), nil
	default:
		return nil, fmt.Errorf(
			"db.Config: unknown type \"%s\"; must be %s or %s",
			dbType,
			TypeInMemory,
			TypePostgres,
		)
	}
}
