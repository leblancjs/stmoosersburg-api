package user

import (
	"testing"

	"github.com/leblancjs/stmoosersburg-api/db"
)

func TestRepositoryFactory(t *testing.T) {
	t.Run("returns an in memory repository when passed an in memory database", func(t *testing.T) {
		repo, _ := NewRepository(&db.InMemory{})

		if _, ok := repo.(*inMemoryRepository); !ok {
			t.Fail()
		}
	})

	t.Run("returns a Postgres repository when passed a Postgres database", func(t *testing.T) {
		repo, _ := NewRepository(&db.Postgres{})

		if _, ok := repo.(*postgresRepository); !ok {
			t.Fail()
		}
	})

	t.Run("fails when no repository exists for the given database", func(t *testing.T) {
		if _, err := NewRepository(nil); err == nil {
			t.Fail()
		}
	})
}
