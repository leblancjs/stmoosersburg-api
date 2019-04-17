package user

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMakingHandler(t *testing.T) {
	t.Run("returns a handler when all is well", func(t *testing.T) {
		handler := MakeHandler(&mockService{})
		if handler == nil {
			t.Fail()
		}
	})
}

func TestDecodingRegisterUserRequest(t *testing.T) {
	username := "Moose"
	email := "moose@stmoosersburg.com"
	password := "P@ssw0rd"

	t.Run("fails when JSON decoder fails", func(t *testing.T) {
		httpReq, _ := http.NewRequest(
			"POST",
			"/users",
			bytes.NewBuffer([]byte("not.json.at.all")),
		)

		if _, err := decodeRegisterUserRequest(nil, httpReq); err == nil {
			t.Fail()
		}
	})

	t.Run("returns a register user request when all is well", func(t *testing.T) {
		httpReq, _ := http.NewRequest(
			"POST",
			"/users",
			bytes.NewBuffer([]byte(fmt.Sprintf(
				`{
					"username": "%s",
					"email": "%s",
					"password": "%s"
				}`,
				username,
				email,
				password,
			))),
		)

		req, err := decodeRegisterUserRequest(nil, httpReq)
		if err != nil {
			t.Fail()
		}
		if req == nil {
			t.FailNow()
		}

		regUserReq, ok := req.(registerUserRequest)
		if !ok {
			t.FailNow()
		}
		if strings.Compare(username, regUserReq.Username) != 0 {
			t.Fail()
		}
		if strings.Compare(email, regUserReq.Email) != 0 {
			t.Fail()
		}
		if strings.Compare(password, regUserReq.Password) != 0 {
			t.Fail()
		}
	})
}

func TestEncodingRegisterUserResponse(t *testing.T) {
	t.Run("writes HTTP status created when all is well", func(t *testing.T) {
		rr := httptest.NewRecorder()

		response := registerUserResponse{}

		if err := encodeRegisterUserResponse(nil, rr, response); err != nil {
			t.Fail()
		}

		if rr.Code != http.StatusCreated {
			t.Fail()
		}
		if rr.Body.Len() <= 0 {
			t.Fail()
		}
	})
}

func TestEncodingResponse(t *testing.T) {
	t.Run("sets Content-Type header to application/json with UTF-8 charset", func(t *testing.T) {
		rr := httptest.NewRecorder()

		if err := encodeResponse(nil, rr, nil); err != nil {
			t.Fail()
		}

		if strings.Compare("application/json; charset=utf-8", rr.Header().Get("Content-Type")) != 0 {
			t.Fail()
		}
	})

	t.Run("writes HTTP status OK when all is well", func(t *testing.T) {
		rr := httptest.NewRecorder()

		if err := encodeResponse(nil, rr, nil); err != nil {
			t.Fail()
		}

		if rr.Code != http.StatusOK {
			t.Fail()
		}
	})
}

func TestEncodingError(t *testing.T) {
	t.Run("writes HTTP status internal server error by default", func(t *testing.T) {
		rr := httptest.NewRecorder()

		encodeError(nil, rr, fmt.Errorf("a terrible error"))

		if rr.Code != http.StatusInternalServerError {
			t.Fail()
		}
	})
}
