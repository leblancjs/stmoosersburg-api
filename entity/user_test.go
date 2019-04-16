package entity

import (
	"strings"
	"testing"
)

var user = User{
	ID:       "a.very.special.moose",
	Username: "Moose",
	Email:    "moose@stmoosersburg.com",
	Password: "a.terrible.password.never.to.be.used",
}

func TestUserValidation(t *testing.T) {
	t.Run("fails when username is missing", func(t *testing.T) {
		u := user
		u.Username = ""

		if err := u.Validate(); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when email is missing", func(t *testing.T) {
		u := user
		u.Email = ""

		if err := u.Validate(); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when password is missing", func(t *testing.T) {
		u := user
		u.Password = ""

		if err := u.Validate(); err == nil {
			t.Fail()
		}
	})

	t.Run("returns nil when all is well", func(t *testing.T) {
		u := user

		if err := u.Validate(); err != nil {
			t.Fail()
		}
	})
}

func TestUserStringFormat(t *testing.T) {
	t.Run("never prints password", func(t *testing.T) {
		userString := user.String()

		if strings.Contains(userString, user.Password) {
			t.Fail()
		}
	})
}
