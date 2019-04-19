package user

import (
	"database/sql"
	"fmt"

	"github.com/leblancjs/stmoosersburg-api/db"
	"github.com/leblancjs/stmoosersburg-api/entity"
)

const (
	createQueryFormat = "INSERT INTO users(username, email, password) VALUES('%s', '%s', '%s') RETURNING id"
	getByIDQuery      = "SELECT id, username, email, password FROM users WHERE id = $1"
	getByEmailQuery   = "SELECT id, username, email, password FROM users WHERE email = $1"
)

type postgresRepository struct {
	database *db.Postgres
}

func NewPostgresRepository(database *db.Postgres) Repository {
	return &postgresRepository{database}
}

func (pr *postgresRepository) Create(username string, email string, password string) (*entity.User, error) {
	user := entity.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	query := fmt.Sprintf(createQueryFormat, username, email, password)

	err := pr.database.QueryRow(query).Scan(&user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(
				"user.PostgresRepository.Create: failed to retrieve user ID",
			)
		}

		return nil, fmt.Errorf(
			"user.PostgresRepository.Create: failed to execute query (%s)",
			err,
		)
	}

	return &user, nil
}

func (pr *postgresRepository) GetByID(id string) (*entity.User, error) {
	var user entity.User

	err := pr.database.QueryRow(getByIDQuery, id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(
				"user.PostgresRepository.GetByID: no user exists with ID \"%s\"",
				id,
			)
		}

		return nil, fmt.Errorf(
			"user.PostgresRepository.GetByID: failed to execute query (%s)",
			err,
		)
	}

	return &user, nil
}

func (pr *postgresRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := pr.database.QueryRow(getByEmailQuery, email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(
				"user.PostgresRepository.GetByEmail: no user exists with email \"%s\"",
				email,
			)
		}

		return nil, fmt.Errorf(
			"user.PostgresRepository.GetByEmail: failed to execute query (%s)",
			err,
		)
	}

	return &user, nil
}
