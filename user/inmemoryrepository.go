package user

import (
	"strconv"

	"github.com/leblancjs/stmoosersburg-api/entity"
)

type inMemoryRepository struct {
	nextID       int
	usersByID    map[string]*entity.User
	usersByEmail map[string]*entity.User
}

func NewInMemoryRepository() Repository {
	return &inMemoryRepository{
		0,
		make(map[string]*entity.User),
		make(map[string]*entity.User),
	}
}

func (repo *inMemoryRepository) Create(username string, email string, password string) (*entity.User, error) {
	user := entity.User{
		ID:       strconv.Itoa(repo.nextID),
		Username: username,
		Email:    email,
		Password: password,
	}

	repo.nextID++

	repo.usersByID[user.ID] = &user
	repo.usersByEmail[user.Email] = &user

	return &user, nil
}

func (repo *inMemoryRepository) GetByID(id string) (*entity.User, error) {
	user, ok := repo.usersByID[id]
	if !ok {
		return nil, nil
	}

	return user, nil
}

func (repo *inMemoryRepository) GetByEmail(email string) (*entity.User, error) {
	user, ok := repo.usersByEmail[email]
	if !ok {
		return nil, nil
	}

	return user, nil
}
