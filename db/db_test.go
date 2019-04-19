package db

import (
	"testing"
)

func TestDatabaseCreation(t *testing.T) {
	t.Run("fails when database type is not recognized", func(t *testing.T) {
		if _, err := New("not.a.type", Config{}); err == nil {
			t.Fail()
		}
	})

	t.Run("returns an in memory database", func(t *testing.T) {
		db, _ := New(TypeInMemory, Config{})

		if _, ok := db.(*InMemory); !ok {
			t.Fail()
		}
	})

	t.Run("returns a Postgres database with the given configuration", func(t *testing.T) {
		db, _ := New(TypePostgres, Config{})

		if _, ok := db.(*Postgres); !ok {
			t.Fail()
		}
	})
}
