package user

import (
	"context"

	"github.com/leblancjs/stmoosersburg-api/endpoint"
)

type registerUserRequest struct {
	Username string
	Email    string
	Password string
}

type registerUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func makeRegisterUserEndpoint(us Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(registerUserRequest)

		u, err := us.Register(req.Username, req.Email, req.Password)
		if err != nil {
			return nil, err
		}

		return &registerUserResponse{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		}, nil
	}
}

type getUserByIDRequest struct {
	ID string
}

type getUserByIDResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func makeGetUserByIDEndpoint(us Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserByIDRequest)

		u, err := us.GetByID(req.ID)
		if err != nil {
			return nil, err
		}

		return &getUserByIDResponse{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		}, nil
	}
}
