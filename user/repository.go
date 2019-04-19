package user

import (
	"fmt"

	"github.com/leblancjs/stmoosersburg-api/db"
	"github.com/leblancjs/stmoosersburg-api/entity"
)

type Repository interface {
	Create(username string, email string, password string) (*entity.User, error)
	GetByID(id string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
}

func NewRepository(database db.DB) (Repository, error) {
	if inmemory, ok := database.(*db.InMemory); ok {
		return NewInMemoryRepository(inmemory), nil
	} else if postgres, ok := database.(*db.Postgres); ok {
		return NewPostgresRepository(postgres), nil
	}

	return nil, fmt.Errorf("user.NewRepository: unsupported database type")
}
