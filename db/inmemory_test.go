package db

import "testing"

func TestOpeningInMemoryDatabase(t *testing.T) {
	t.Run("creates an empty array of users when all is well", func(t *testing.T) {
		db := InMemory{}

		if err := db.Open(); err != nil {
			t.Fail()
		}

		if db.Users == nil {
			t.FailNow()
		}

		if len(db.Users) != 0 {
			t.Fail()
		}
	})
}

func TestClosingInMemoryDatabase(t *testing.T) {
	t.Run("does nothing when all is well", func(t *testing.T) {
		db := InMemory{}

		if err := db.Close(); err != nil {
			t.Fail()
		}
	})
}
