package user

import (
	"strconv"
	"strings"
	"testing"

	"github.com/leblancjs/stmoosersburg-api/db"
	"github.com/leblancjs/stmoosersburg-api/entity"
)

const (
	id       = "0"
	username = "Moose"
	email    = "moose@stmoosersburg.com"
	password = "a.super.secret.hashed.password"
)

func TestInMemoryRepositoryConstructor(t *testing.T) {
	database := &db.InMemory{}

	t.Run("starts nextID at zero", func(t *testing.T) {
		repo, _ := NewInMemoryRepository(database).(*inMemoryRepository)

		if repo.nextID != 0 {
			t.Fail()
		}
	})

	t.Run("returns a repository with the given database", func(t *testing.T) {
		repo, _ := NewInMemoryRepository(database).(*inMemoryRepository)

		if repo == nil {
			t.Fail()
		}
		if repo.database != database {
			t.Fail()
		}
	})
}

func TestInMemoryRepositoryCreation(t *testing.T) {
	database := &db.InMemory{}
	database.Open()

	t.Run("creates a new user with the given username, email, and password, and a new ID", func(t *testing.T) {
		repo, _ := NewInMemoryRepository(database).(*inMemoryRepository)

		expectedUserID := strconv.Itoa(repo.nextID)

		user, _ := repo.Create(username, email, password)

		if strings.Compare(expectedUserID, user.ID) != 0 {
			t.Fail()
		}
		if strings.Compare(username, user.Username) != 0 {
			t.Fail()
		}
		if strings.Compare(email, user.Email) != 0 {
			t.Fail()
		}
		if strings.Compare(password, user.Password) != 0 {
			t.Fail()
		}
	})

	t.Run("increments nextID", func(t *testing.T) {
		repo, _ := NewInMemoryRepository(database).(*inMemoryRepository)

		initialNextID := repo.nextID

		_, _ = repo.Create(username, email, password)

		if initialNextID == repo.nextID {
			t.Fail()
		}
	})

	t.Run("adds a new user to the database's list of users", func(t *testing.T) {
		repo, _ := NewInMemoryRepository(database).(*inMemoryRepository)

		expectedUser, _ := repo.Create(username, email, password)

		user := repo.database.Users[0]

		if strings.Compare(expectedUser.ID, user.ID) != 0 {
			t.Fail()
		}
	})
}

func TestInMemoryRepositoryGettingByID(t *testing.T) {
	user := entity.User{
		ID:       id,
		Username: username,
		Email:    email,
		Password: password,
	}

	database := &db.InMemory{}
	database.Open()
	database.Users = append(database.Users, user)

	repo := inMemoryRepository{
		nextID:   1,
		database: database,
	}

	t.Run("returns error when no user is found", func(t *testing.T) {
		if _, err := repo.GetByID("no.way.this.exists"); err == nil {
			t.Fail()
		}
	})

	t.Run("returns user with given ID", func(t *testing.T) {
		user, _ := repo.GetByID(id)
		if user == nil {
			t.FailNow()
		}
		if strings.Compare(id, user.ID) != 0 {
			t.Fail()
		}
	})
}

func TestInMemoryRepositoryGettingByEmail(t *testing.T) {
	user := entity.User{
		ID:       id,
		Username: username,
		Email:    email,
		Password: password,
	}

	database := &db.InMemory{}
	database.Open()
	database.Users = append(database.Users, user)

	repo := inMemoryRepository{
		nextID:   1,
		database: database,
	}

	t.Run("returns error when no user is found", func(t *testing.T) {
		if _, err := repo.GetByEmail("no.way.this.exists"); err == nil {
			t.Fail()
		}
	})

	t.Run("returns user with given email", func(t *testing.T) {
		user, _ := repo.GetByEmail(email)
		if user == nil {
			t.FailNow()
		}
		if strings.Compare(email, user.Email) != 0 {
			t.Fail()
		}
	})
}
