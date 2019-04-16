package user

import "github.com/leblancjs/stmoosersburg-api/entity"

type Repository interface {
	Create(username string, email string, password string) (*entity.User, error)
	GetByID(id string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
}
