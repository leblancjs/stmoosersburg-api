package entity

import (
	"fmt"
)

type User struct {
	ID       string
	Username string
	Email    string
	Password string
}

func (u User) Validate() error {
	if u.Username == "" {
		return fmt.Errorf("entity.User.Validate: username is required")
	}

	if u.Email == "" {
		return fmt.Errorf("entity.User.Validate: email is required")
	}

	if u.Password == "" {
		return fmt.Errorf("entity.User.Validate: password is required")
	}

	return nil
}

func (u User) String() string {
	return fmt.Sprintf(
		"User { ID: %s, Username: %s, Email: %s }",
		u.ID,
		u.Username,
		u.Email,
	)
}
