package user

import (
	"testing"

	"github.com/leblancjs/stmoosersburg-api/db"
)

func TestPostgresRepositoryCreation(t *testing.T) {
	database := &db.Postgres{}

	t.Run("returns a postgres repository", func(t *testing.T) {
		if _, ok := NewPostgresRepository(database).(*postgresRepository); !ok {
			t.Fail()
		}
	})

	t.Run("returns a postgres repository that uses the given database", func(t *testing.T) {
		pr := NewPostgresRepository(database).(*postgresRepository)

		if pr.database != database {
			t.Fail()
		}
	})
}
