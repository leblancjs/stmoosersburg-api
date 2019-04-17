package http

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/leblancjs/stmoosersburg-api/endpoint"
)

func TestHandlerConstruction(t *testing.T) {
	mockEndpoint := newMockEndpoint(nil, nil)
	mockRequestDecoder := newMockRequestDecoder(nil, nil)
	mockResponseEncoder := newMockResponseEncoder(nil)
	mockErrorEncoder := newMockErrorEncoder()

	t.Run("fails when endpoint is missing", func(t *testing.T) {
		if _, err := NewHandler(
			nil,
			mockRequestDecoder.decode,
			mockResponseEncoder.encode,
			mockErrorEncoder.encode,
		); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when decode request func is missing", func(t *testing.T) {
		if _, err := NewHandler(
			mockEndpoint.endpoint,
			nil,
			mockResponseEncoder.encode,
			mockErrorEncoder.encode,
		); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when encode response func is missing", func(t *testing.T) {
		if _, err := NewHandler(
			mockEndpoint.endpoint,
			mockRequestDecoder.decode,
			nil,
			mockErrorEncoder.encode,
		); err == nil {
			t.Fail()
		}
	})

	t.Run("fails when encode error func is missing", func(t *testing.T) {
		if _, err := NewHandler(
			mockEndpoint.endpoint,
			mockRequestDecoder.decode,
			mockResponseEncoder.encode,
			nil,
		); err == nil {
			t.Fail()
		}
	})

	t.Run("returns handler when all is well", func(t *testing.T) {
		handler, err := NewHandler(
			mockEndpoint.endpoint,
			mockRequestDecoder.decode,
			mockResponseEncoder.encode,
			mockErrorEncoder.encode,
		)
		if err != nil {
			t.Fail()
		}
		if handler == nil {
			t.Fail()
		}
	})
}

func TestHandlerServing(t *testing.T) {
	t.Run("encodes error and returns when decoding request fails", func(t *testing.T) {
		mockEndpoint := newMockEndpoint(nil, nil)
		mockResponseEncoder := newMockResponseEncoder(nil)
		mockErrorEncoder := newMockErrorEncoder()

		mockRequestDecoder := newMockRequestDecoder(
			nil,
			fmt.Errorf("failed to decode request"),
		)

		handler, _ := NewHandler(
			mockEndpoint.endpoint,
			mockRequestDecoder.decode,
			mockResponseEncoder.encode,
			mockErrorEncoder.encode,
		)

		handler.ServeHTTP(nil, &http.Request{})

		if !mockRequestDecoder.WasCalled() {
			t.Fail()
		}
		if !mockErrorEncoder.WasCalled() {
			t.Fail()
		}
		if mockEndpoint.WasCalled() {
			t.Fail()
		}
		if mockResponseEncoder.WasCalled() {
			t.Fail()
		}
	})

	t.Run("encodes error and returns when endpoint fails", func(t *testing.T) {
		mockRequestDecoder := newMockRequestDecoder(nil, nil)
		mockResponseEncoder := newMockResponseEncoder(nil)
		mockErrorEncoder := newMockErrorEncoder()

		mockEndpoint := newMockEndpoint(
			nil,
			fmt.Errorf("endpoint failed to process request"),
		)

		handler, _ := NewHandler(
			mockEndpoint.endpoint,
			mockRequestDecoder.decode,
			mockResponseEncoder.encode,
			mockErrorEncoder.encode,
		)

		handler.ServeHTTP(nil, &http.Request{})

		if !mockRequestDecoder.WasCalled() {
			t.Fail()
		}
		if !mockEndpoint.WasCalled() {
			t.Fail()
		}
		if !mockErrorEncoder.WasCalled() {
			t.Fail()
		}
		if mockResponseEncoder.WasCalled() {
			t.Fail()
		}
	})

	t.Run("encodes error and returns when encoding response fails", func(t *testing.T) {
		mockEndpoint := newMockEndpoint(nil, nil)
		mockRequestDecoder := newMockRequestDecoder(nil, nil)
		mockErrorEncoder := newMockErrorEncoder()

		mockResponseEncoder := newMockResponseEncoder(
			fmt.Errorf("failed to encode response"),
		)

		handler, _ := NewHandler(
			mockEndpoint.endpoint,
			mockRequestDecoder.decode,
			mockResponseEncoder.encode,
			mockErrorEncoder.encode,
		)

		handler.ServeHTTP(nil, &http.Request{})

		if !mockRequestDecoder.WasCalled() {
			t.Fail()
		}
		if !mockEndpoint.WasCalled() {
			t.Fail()
		}
		if !mockResponseEncoder.WasCalled() {
			t.Fail()
		}
		if !mockErrorEncoder.WasCalled() {
			t.Fail()
		}
	})
}

type mockFunc struct {
	callCount int
}

func (mf *mockFunc) WasCalled() bool {
	return mf.callCount > 0
}

type mockEndpoint struct {
	mockFunc
	endpoint endpoint.Endpoint
}

func newMockEndpoint(response interface{}, err error) *mockEndpoint {
	mock := mockEndpoint{}
	mock.endpoint = func(_ context.Context, _ interface{}) (interface{}, error) {
		mock.callCount++
		return response, err
	}
	return &mock
}

type mockRequestDecoder struct {
	mockFunc
	decode DecodeRequestFunc
}

func newMockRequestDecoder(request interface{}, err error) *mockRequestDecoder {
	mock := mockRequestDecoder{}
	mock.decode = func(_ context.Context, _ *http.Request) (interface{}, error) {
		mock.callCount++
		return request, err
	}
	return &mock
}

type mockResponseEncoder struct {
	mockFunc
	encode EncodeResponseFunc
}

func newMockResponseEncoder(err error) *mockResponseEncoder {
	mock := mockResponseEncoder{}
	mock.encode = func(_ context.Context, _ http.ResponseWriter, response interface{}) error {
		mock.callCount++
		return err
	}
	return &mock
}

type mockErrorEncoder struct {
	mockFunc
	encode EncodeErrorFunc
}

func newMockErrorEncoder() *mockErrorEncoder {
	mock := mockErrorEncoder{}
	mock.encode = func(_ context.Context, _ http.ResponseWriter, err error) {
		mock.callCount++
	}
	return &mock
}
