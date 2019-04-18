package db

import (
	"strings"
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
		if db == nil {
			t.FailNow()
		}

		_, ok := db.(*InMemory)
		if !ok {
			t.FailNow()
		}
	})

	t.Run("returns a Postgres database with the given configuration", func(t *testing.T) {
		conf := Config{
			Host:     "host",
			Port:     "1234",
			User:     "user",
			Password: "password",
			Name:     "name",
		}

		db, _ := New(TypePostgres, conf)
		if db == nil {
			t.FailNow()
		}

		postgresDB, ok := db.(*Postgres)
		if !ok {
			t.FailNow()
		}

		if strings.Compare(conf.Host, postgresDB.conf.Host) != 0 {
			t.Fail()
		}
		if strings.Compare(conf.Port, postgresDB.conf.Port) != 0 {
			t.Fail()
		}
		if strings.Compare(conf.User, postgresDB.conf.User) != 0 {
			t.Fail()
		}
		if strings.Compare(conf.Password, postgresDB.conf.Password) != 0 {
			t.Fail()
		}
		if strings.Compare(conf.Name, postgresDB.conf.Name) != 0 {
			t.Fail()
		}
	})
}
