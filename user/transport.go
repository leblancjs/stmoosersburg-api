package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	stmhttp "github.com/leblancjs/stmoosersburg-api/transport/http"
)

func MakeHandler(us Service) http.Handler {
	registerUserHandler := stmhttp.NewHandler(
		makeRegisterUserEndpoint(us),
		decodeRegisterUserRequest,
		encodeRegisterUserResponse,
		encodeError,
	)

	getUserByIDHandler := stmhttp.NewHandler(
		makeGetUserByIDEndpoint(us),
		decodeGetUserByIDRequest,
		encodeResponse,
		encodeError,
	)

	r := mux.NewRouter()

	r.Handle("/v1/users", registerUserHandler).Methods("POST")
	r.Handle("/v1/users/{id}", getUserByIDHandler).Methods("GET")

	return r
}

func decodeRegisterUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return registerUserRequest{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	}, nil
}

func encodeRegisterUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.WriteHeader(http.StatusCreated)
	return encodeResponse(ctx, w, response)
}

func decodeGetUserByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		return nil, fmt.Errorf("bad route")
	}

	return getUserByIDRequest{
		ID: id,
	}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(ctx context.Context, w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
