package user

import (
	"fmt"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

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

func TestPostgresRepositoryCreatingUser(t *testing.T) {
	queryResultColumns := []string{"id"}
	expectedQuery := fmt.Sprintf(createQueryFormat, mockUserUsername, mockUserEmail, mockUserPassword)

	t.Run("fails when query returns no rows (user ID can't be retrieved)", func(t *testing.T) {
		database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("failed to open mock database connection (%s)", err)
		}
		defer database.Close()

		pr := NewPostgresRepository(&db.Postgres{DB: database})

		mock.ExpectQuery(expectedQuery).
			WillReturnRows(mock.NewRows(queryResultColumns))

		if _, err := pr.Create(mockUserUsername, mockUserEmail, mockUserPassword); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when query fails", func(t *testing.T) {
		database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("failed to open mock database connection (%s)", err)
		}
		defer database.Close()

		pr := NewPostgresRepository(&db.Postgres{DB: database})

		mock.ExpectQuery(expectedQuery).
			WillReturnError(fmt.Errorf("an error occurred"))

		if _, err := pr.Create(mockUserUsername, mockUserEmail, mockUserPassword); err == nil {
			t.Fail()
		}
	})

	t.Run("returns the created user with its new ID when all is well", func(t *testing.T) {
		database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("failed to open mock database connection (%s)", err)
		}
		defer database.Close()

		pr := NewPostgresRepository(&db.Postgres{DB: database})

		mock.ExpectQuery(expectedQuery).
			WillReturnRows(sqlmock.NewRows(queryResultColumns).AddRow(mockUserID))

		user, err := pr.Create(mockUserUsername, mockUserEmail, mockUserPassword)
		if err != nil {
			t.FailNow()
		}
		if strings.Compare(mockUserID, user.ID) != 0 {
			t.Fail()
		}
		if strings.Compare(mockUserUsername, user.Username) != 0 {
			t.Fail()
		}
		if strings.Compare(mockUserEmail, user.Email) != 0 {
			t.Fail()
		}
		if strings.Compare(mockUserPassword, user.Password) != 0 {
			t.Fail()
		}
	})
}

func TestPostgresRepositoryGettingUserByID(t *testing.T) {
	queryResultColumns := []string{"id", "username", "email", "password"}
	expectedQuery := getByIDQuery

	t.Run("fails when query returns no rows (no user with given ID exists)", func(t *testing.T) {
		database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("failed to open mock database connection (%s)", err)
		}
		defer database.Close()

		pr := NewPostgresRepository(&db.Postgres{DB: database})

		mock.ExpectQuery(expectedQuery).
			WithArgs(mockUserID).
			WillReturnRows(mock.NewRows(queryResultColumns))

		if _, err := pr.GetByID(mockUserID); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when query fails", func(t *testing.T) {
		database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("failed to open mock database connection (%s)", err)
		}
		defer database.Close()

		pr := NewPostgresRepository(&db.Postgres{DB: database})

		mock.ExpectQuery(expectedQuery).
			WillReturnError(fmt.Errorf("an error occurred"))

		if _, err := pr.GetByID(mockUserID); err == nil {
			t.Fail()
		}
	})

	t.Run("returns the user with the given ID when all is well", func(t *testing.T) {
		database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("failed to open mock database connection (%s)", err)
		}
		defer database.Close()

		pr := NewPostgresRepository(&db.Postgres{DB: database})

		mock.ExpectQuery(expectedQuery).
			WithArgs(mockUserID).
			WillReturnRows(
				sqlmock.NewRows(queryResultColumns).
					AddRow(mockUserID, mockUserUsername, mockUserEmail, mockUserPassword),
			)

		user, err := pr.GetByID(mockUserID)
		if err != nil {
			t.FailNow()
		}
		if strings.Compare(mockUserID, user.ID) != 0 {
			t.Fail()
		}
		if strings.Compare(mockUserUsername, user.Username) != 0 {
			t.Fail()
		}
		if strings.Compare(mockUserEmail, user.Email) != 0 {
			t.Fail()
		}
		if strings.Compare(mockUserPassword, user.Password) != 0 {
			t.Fail()
		}
	})
}

func TestPostgresRepositoryGettingUserByEmail(t *testing.T) {
	queryResultColumns := []string{"id", "username", "email", "password"}
	expectedQuery := getByEmailQuery

	t.Run("fails when query returns no rows (no user with given email exists)", func(t *testing.T) {
		database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("failed to open mock database connection (%s)", err)
		}
		defer database.Close()

		pr := NewPostgresRepository(&db.Postgres{DB: database})

		mock.ExpectQuery(expectedQuery).
			WithArgs(mockUserEmail).
			WillReturnRows(mock.NewRows(queryResultColumns))

		if _, err := pr.GetByEmail(mockUserEmail); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when query fails", func(t *testing.T) {
		database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("failed to open mock database connection (%s)", err)
		}
		defer database.Close()

		pr := NewPostgresRepository(&db.Postgres{DB: database})

		mock.ExpectQuery(expectedQuery).
			WillReturnError(fmt.Errorf("an error occurred"))

		if _, err := pr.GetByEmail(mockUserEmail); err == nil {
			t.Fail()
		}
	})

	t.Run("returns the user with the given email when all is well", func(t *testing.T) {
		database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("failed to open mock database connection (%s)", err)
		}
		defer database.Close()

		pr := NewPostgresRepository(&db.Postgres{DB: database})

		mock.ExpectQuery(expectedQuery).
			WithArgs(mockUserEmail).
			WillReturnRows(
				sqlmock.NewRows(queryResultColumns).
					AddRow(mockUserID, mockUserUsername, mockUserEmail, mockUserPassword),
			)

		user, err := pr.GetByEmail(mockUserEmail)
		if err != nil {
			t.FailNow()
		}
		if strings.Compare(mockUserID, user.ID) != 0 {
			t.Fail()
		}
		if strings.Compare(mockUserUsername, user.Username) != 0 {
			t.Fail()
		}
		if strings.Compare(mockUserEmail, user.Email) != 0 {
			t.Fail()
		}
		if strings.Compare(mockUserPassword, user.Password) != 0 {
			t.Fail()
		}
	})
}
