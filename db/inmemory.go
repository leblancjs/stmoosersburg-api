package db

import "github.com/leblancjs/stmoosersburg-api/entity"

// InMemory represents an in memory database.
//
// All entities are persisted in memory, so they are lost when the service is
// terminated.
type InMemory struct {
	db
	Users []entity.User
}

// Open opens the in memory database by creating the appropriate collections.
func (db *InMemory) Open() error {
	db.Users = make([]entity.User, 0)

	return nil
}

// Close closes the in memory database by doing absolutely nothing.
func (db *InMemory) Close() error {
	return nil
}
