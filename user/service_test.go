package user

import (
	"fmt"
	"testing"

	"github.com/leblancjs/stmoosersburg-api/entity"
)

func TestServiceConstructor(t *testing.T) {
	repo := &mockRepository{}
	hashSvc := &mockHashService{}

	t.Run("fails when repository is missing", func(t *testing.T) {
		if _, err := NewService(nil, hashSvc); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when hash service is missing", func(t *testing.T) {
		if _, err := NewService(repo, nil); err == nil {
			t.Fail()
		}
	})

	t.Run("returns a service with repo and hash service", func(t *testing.T) {
		svc, _ := NewService(repo, hashSvc)
		if svc == nil {
			t.FailNow()
		}

		userSvc, _ := svc.(*service)
		if userSvc.repo != repo {
			t.Fail()
		}
		if userSvc.hashSvc != hashSvc {
			t.Fail()
		}
	})
}

func TestServiceRegistration(t *testing.T) {
	username := "moose"
	email := "moose@stmoosersburg.com"
	password := "P@ssw0rd"

	t.Run("fails when username validation fails", func(t *testing.T) {
		svc, _ := NewService(&mockRepository{}, &mockHashService{})

		if _, err := svc.Register("", email, password); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when email validation fails", func(t *testing.T) {
		svc, _ := NewService(&mockRepository{}, &mockHashService{})

		if _, err := svc.Register(username, "", password); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when password validation fails", func(t *testing.T) {
		svc, _ := NewService(&mockRepository{}, &mockHashService{})

		if _, err := svc.Register(username, email, ""); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when hash generation fails", func(t *testing.T) {
		svc, _ := NewService(&mockRepository{}, &mockHashService{true, false})

		if _, err := svc.Register(username, email, password); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when creation in repository fails", func(t *testing.T) {
		svc, _ := NewService(&mockRepository{failOnCreate: true}, &mockHashService{})

		if _, err := svc.Register(username, email, password); err == nil {
			t.Fail()
		}
	})

	t.Run("returns new user when all is well", func(t *testing.T) {
		svc, _ := NewService(&mockRepository{}, &mockHashService{})

		user, err := svc.Register(username, email, password)
		if err != nil {
			t.Fail()
		}
		if user == nil {
			t.Fail()
		}
	})
}

func TestServiceGettingByID(t *testing.T) {
	id := "a.very.unique.identifier"

	t.Run("fails when getting from repository fails", func(t *testing.T) {
		svc, _ := NewService(&mockRepository{failOnGetByID: true}, &mockHashService{})

		if _, err := svc.GetByID(id); err == nil {
			t.Fail()
		}
	})

	t.Run("returns user when all is well", func(t *testing.T) {
		svc, _ := NewService(&mockRepository{}, &mockHashService{})

		user, err := svc.GetByID(id)
		if err != nil {
			t.Fail()
		}
		if user == nil {
			t.Fail()
		}
	})
}

func TestServiceGettingByEmail(t *testing.T) {
	email := "moose@stmoosersburg.com"

	t.Run("fails when getting from repository fails", func(t *testing.T) {
		svc, _ := NewService(&mockRepository{failOnGetByEmail: true}, &mockHashService{})

		if _, err := svc.GetByEmail(email); err == nil {
			t.Fail()
		}
	})

	t.Run("returns user when all is well", func(t *testing.T) {
		svc, _ := NewService(&mockRepository{}, &mockHashService{})

		user, err := svc.GetByEmail(email)
		if err != nil {
			t.Fail()
		}
		if user == nil {
			t.Fail()
		}
	})
}

func TestServiceUsernameValidation(t *testing.T) {
	t.Run("fails when username is empty", func(t *testing.T) {
		if err := validateUsername(""); err == nil {
			t.Fail()
		}
	})

	t.Run("succeeds when username is not empty", func(t *testing.T) {
		if err := validateUsername("not.empty"); err != nil {
			t.Fail()
		}
	})
}

func TestServiceEmailValidation(t *testing.T) {
	// These invalid and valid email formats were taken from the following page:
	// https://help.xmatters.com/ondemand/trial/valid_email_format.htm.
	//
	// It is not meant to be an exhaustive test... but it helps to check more
	// than one.
	invalidEmails := []string{
		"abc-@mail.com",
		"abc..def@mail.com",
		".abc@mail.com",
		"-abc@mail.com",
		"abc#def@mail.com",
		"abc.def@mail.c",
		"abc.def@mail#archive.com",
		"abc.def@mail",
		"abc.def@mail..com",
	}

	validEmails := []string{
		"abc-d@mail.com",
		"abc.def@mail.com",
		"abc@mail.com",
		"abc_def@mail.com",
		"123-d@mail.com",
		"123.456@mail.com",
		"123@mail.com",
		"123_456@mail.com",
		"abc.def@mail.cc",
		"abc.def@mail-archive.com",
		"abc.def@mail.org",
		"abc.def@mail.com",
		"abc.def@123.com",
		"abc.def@123-456.com",
	}

	t.Run("fails when email is empty", func(t *testing.T) {
		if err := validateEmail(""); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when email has invalid format", func(t *testing.T) {
		for _, invalidEmail := range invalidEmails {
			if err := validateEmail(invalidEmail); err == nil {
				t.Fail()
			}
		}
	})

	t.Run("succeeds when email has valid format", func(t *testing.T) {
		for _, validEmail := range validEmails {
			if err := validateEmail(validEmail); err != nil {
				t.Fail()
			}
		}
	})
}

func TestServicePasswordValidation(t *testing.T) {
	validPassword := "P@ssw0rd"
	passwordWithNoLowerCaseLetter := "NOLOWERCASE@1"
	passwordWithNoUpperCaseLetter := "nouppercase@1"
	passwordWithNoDigit := "Nodigit@all"
	passwordWithNoSpecialCharacter := "Nospecialcharacter1"
	passwordTooShort := "To0b@d"

	t.Run("fails when password is empty", func(t *testing.T) {
		if err := validatePassword(""); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when password does not have at least one lower case letter", func(t *testing.T) {
		if err := validatePassword(passwordWithNoLowerCaseLetter); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when password does not have at least one upper case letter", func(t *testing.T) {
		if err := validatePassword(passwordWithNoUpperCaseLetter); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when password does not have at least one digit", func(t *testing.T) {
		if err := validatePassword(passwordWithNoDigit); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when password does not have at least one special character", func(t *testing.T) {
		if err := validatePassword(passwordWithNoSpecialCharacter); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when password has fewer than 8 characters", func(t *testing.T) {
		if err := validatePassword(passwordTooShort); err == nil {
			t.Fail()
		}
	})

	t.Run("succeeds when password has valid format", func(t *testing.T) {
		if err := validatePassword(validPassword); err != nil {
			t.Fail()
		}
	})
}

type mockRepository struct {
	failOnCreate     bool
	failOnGetByID    bool
	failOnGetByEmail bool
}

func (mock *mockRepository) Create(username, email, password string) (*entity.User, error) {
	if mock.failOnCreate {
		return nil, fmt.Errorf("failed to create user")
	}

	return &entity.User{
		ID:       "mock",
		Username: username,
		Email:    email,
		Password: password,
	}, nil
}

func (mock *mockRepository) GetByID(id string) (*entity.User, error) {
	if mock.failOnGetByID {
		return nil, fmt.Errorf("failed to get user by ID")
	}

	return &entity.User{
		ID:       id,
		Username: "username",
		Email:    "email@address.com",
		Password: "P@ssw0rd",
	}, nil
}

func (mock *mockRepository) GetByEmail(email string) (*entity.User, error) {
	if mock.failOnGetByEmail {
		return nil, fmt.Errorf("failed to get user by email")
	}

	return &entity.User{
		ID:       id,
		Username: "username",
		Email:    email,
		Password: "P@ssw0rd",
	}, nil
}

type mockHashService struct {
	failOnHashGeneration bool
	failOnHashComparison bool
}

func (mock *mockHashService) GenerateFromPassword(password string) (string, error) {
	if mock.failOnHashGeneration {
		return "", fmt.Errorf("failed to generate hash from password")
	}

	return password, nil
}

func (mock *mockHashService) MatchPassword(hash, password string) bool {
	return mock.failOnHashComparison
}
