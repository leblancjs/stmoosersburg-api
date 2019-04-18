package user

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leblancjs/stmoosersburg-api/db"
	"github.com/leblancjs/stmoosersburg-api/entity"
)

type inMemoryRepository struct {
	nextID   int
	database *db.InMemory
}

func NewInMemoryRepository(database *db.InMemory) Repository {
	return &inMemoryRepository{0, database}
}

func (repo *inMemoryRepository) Create(username string, email string, password string) (*entity.User, error) {
	user := entity.User{
		ID:       strconv.Itoa(repo.nextID),
		Username: username,
		Email:    email,
		Password: password,
	}

	repo.nextID++

	repo.database.Users = append(repo.database.Users, user)

	return &user, nil
}

func (repo *inMemoryRepository) GetByID(id string) (*entity.User, error) {
	var user *entity.User

	for _, u := range repo.database.Users {
		if strings.Compare(id, u.ID) == 0 {
			user = &u
			break
		}
	}

	if user == nil {
		return nil, fmt.Errorf("user.InMemoryRepository.GetByID: no user exists with ID \"%s\"", id)
	}

	return user, nil
}

func (repo *inMemoryRepository) GetByEmail(email string) (*entity.User, error) {
	var user *entity.User

	for _, u := range repo.database.Users {
		if strings.Compare(email, u.Email) == 0 {
			user = &u
			break
		}
	}

	if user == nil {
		return nil, fmt.Errorf("user.InMemoryRepository.GetByEmail: no user exists with email \"%s\"", email)
	}

	return user, nil
}
