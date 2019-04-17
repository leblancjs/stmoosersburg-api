package user

import (
	"fmt"
	"regexp"

	"github.com/leblancjs/stmoosersburg-api/entity"
	"github.com/leblancjs/stmoosersburg-api/hash"
)

type Service interface {
	Register(username string, email string, password string) (*entity.User, error)
	GetByID(id string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
}

type service struct {
	repo    Repository
	hashSvc hash.Service
}

func NewService(repo Repository, hashSvc hash.Service) (Service, error) {
	if repo == nil {
		return nil, fmt.Errorf("user.NewService: repository is required")
	}

	if hashSvc == nil {
		return nil, fmt.Errorf("user.NewService: hash service is required")
	}

	return &service{
		repo,
		hashSvc,
	}, nil
}

func (svc *service) Register(username string, email string, password string) (*entity.User, error) {
	if err := validateUsername(username); err != nil {
		return nil, fmt.Errorf("user.Service.Register: %s", err)
	}

	if err := validateEmail(email); err != nil {
		return nil, fmt.Errorf("user.Service.Register: %s", err)
	}

	if err := validatePassword(password); err != nil {
		return nil, fmt.Errorf("user.Service.Register: %s", err)
	}

	if _, err := svc.repo.GetByEmail(email); err == nil {
		return nil, fmt.Errorf("user.Service.Register: user already exists with emal \"%s\"", email)
	}

	hashedPassword, err := svc.hashSvc.GenerateFromPassword(password)
	if err != nil {
		return nil, fmt.Errorf("user.Service.Register: failed to hash password (%s)", err)
	}

	user, err := svc.repo.Create(username, email, hashedPassword)
	if err != nil {
		return nil, fmt.Errorf("user.Service.Register: failed to create user (%s)", err)
	}

	return user, nil
}

func (svc *service) GetByID(id string) (*entity.User, error) {
	user, err := svc.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("user.Service.GetByID: %s", err)
	}

	return user, nil
}

func (svc *service) GetByEmail(email string) (*entity.User, error) {
	user, err := svc.repo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user.Service.GetByEmail: %s", err)
	}

	return user, nil
}

const (
	// Credit for emailRegexp goes to Andy Smith.
	// http://www.regexlib.com/REDetails.aspx?regexp_id=26
	//
	// I modified the prefix verification to make sure that periods, dashes,
	// and underscores are preceded and succeeded with a letter or number.
	emailRegexp = `^(([a-zA-Z0-9]|[a-zA-Z0-9][_\-\.][a-zA-Z0-9])+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$`

	minPasswordLength = 8

	// Credit for passwordRegexp goes to Nic Raboy.
	// https://www.thepolyglotdeveloper.com/2015/05/use-regex-to-test-password-strength-in-javascript/
	//
	// It's been broken down into different expressions, since Go does not
	// support backtracking.
	//
	// The password must be at least 8 characters long, contain at least one
	// lowercase, one uppercase letter, one number, and one special character.
	passwordAtLeastOneLowerCaseLetterRegexp  = `^.*[a-z].*$`
	passwordAtLeastOneUpperCaseLetterRegexp  = `^.*[A-Z].*$`
	passwordAtLeastOneDigitRegexp            = `^.*\d.*$`
	passwordAtLeastOneSpecialCharacterRegexp = `^.*[!@#\$%\^&\*].*$`
)

func validateUsername(username string) error {
	if username == "" {
		return fmt.Errorf("username is required")
	}

	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email is required")
	}

	matched, _ := regexp.MatchString(emailRegexp, email)
	if !matched {
		return fmt.Errorf("email is malformed")
	}

	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("password is required")
	}

	if len(password) < minPasswordLength {
		return fmt.Errorf("password must be at least %d character(s) long", minPasswordLength)
	}
	matched, _ := regexp.MatchString(passwordAtLeastOneLowerCaseLetterRegexp, password)
	if !matched {
		return fmt.Errorf("password is missing lower case letter (a-z)")
	}
	matched, _ = regexp.MatchString(passwordAtLeastOneUpperCaseLetterRegexp, password)
	if !matched {
		return fmt.Errorf("password is missing upper case letter (A-Z)")
	}
	matched, _ = regexp.MatchString(passwordAtLeastOneDigitRegexp, password)
	if !matched {
		return fmt.Errorf("password is missing digit (0-9)")
	}
	matched, _ = regexp.MatchString(passwordAtLeastOneSpecialCharacterRegexp, password)
	if !matched {
		return fmt.Errorf("password is missing special character")
	}

	return nil
}
