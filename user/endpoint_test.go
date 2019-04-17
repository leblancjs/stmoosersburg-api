package user

import (
	"fmt"
	"strings"
	"testing"

	"github.com/leblancjs/stmoosersburg-api/entity"
)

func TestRegisterUserEndpoint(t *testing.T) {
	req := registerUserRequest{
		Username: mockUserUsername,
		Email:    mockUserEmail,
		Password: mockUserPassword,
	}

	t.Run("fails when user service fails", func(t *testing.T) {
		endpoint := makeRegisterUserEndpoint(&mockService{failOnRegister: true})

		if _, err := endpoint(nil, req); err == nil {
			t.Fail()
		}
	})

	t.Run("returns register user response when all is well", func(t *testing.T) {
		endpoint := makeRegisterUserEndpoint(&mockService{})

		resp, _ := endpoint(nil, req)
		if resp == nil {
			t.Fail()
		}

		regUserResp, ok := resp.(*registerUserResponse)
		if !ok {
			t.Fail()
		}

		if strings.Compare(mockUserID, regUserResp.ID) != 0 {
			t.Fail()
		}
		if strings.Compare(req.Username, regUserResp.Username) != 0 {
			t.Fail()
		}
		if strings.Compare(req.Email, regUserResp.Email) != 0 {
			t.Fail()
		}
	})
}

func TestGetUserByIDEndpoint(t *testing.T) {
	req := getUserByIDRequest{
		ID: mockUserID,
	}

	t.Run("fails when user service fails", func(t *testing.T) {
		endpoint := makeGetUserByIDEndpoint(&mockService{failOnGetByID: true})

		if _, err := endpoint(nil, req); err == nil {
			t.Fail()
		}
	})

	t.Run("returns get user by ID response when all is well", func(t *testing.T) {
		endpoint := makeGetUserByIDEndpoint(&mockService{})

		resp, _ := endpoint(nil, req)
		if resp == nil {
			t.Fail()
		}

		getUserResp, ok := resp.(*getUserByIDResponse)
		if !ok {
			t.Fail()
		}

		if strings.Compare(mockUserID, getUserResp.ID) != 0 {
			t.Fail()
		}
	})
}

const (
	mockUserID       = "mock.user.id"
	mockUserUsername = "Moose"
	mockUserEmail    = "moose@stmoosersburg.com"
	mockUserPassword = "P@ssw0rd"
)

type mockService struct {
	failOnRegister   bool
	failOnGetByID    bool
	failOnGetByEmail bool
}

func (mock *mockService) Register(username string, email string, password string) (*entity.User, error) {
	if mock.failOnRegister {
		return nil, fmt.Errorf("failed to register user")
	}

	return &entity.User{
		ID:       mockUserID,
		Username: username,
		Email:    email,
		Password: password,
	}, nil
}

func (mock *mockService) GetByID(id string) (*entity.User, error) {
	if mock.failOnGetByID {
		return nil, fmt.Errorf("failed to get user by ID")
	}

	return &entity.User{ID: id}, nil
}

func (mock *mockService) GetByEmail(email string) (*entity.User, error) {
	if mock.failOnGetByEmail {
		return nil, fmt.Errorf("failed to get user by email")
	}

	return &entity.User{Email: email}, nil
}
