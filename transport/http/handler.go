package http

import (
	"context"
	"net/http"

	"github.com/leblancjs/stmoosersburg-api/endpoint"
)

type DecodeRequestFunc func(ctx context.Context, r *http.Request) (request interface{}, err error)
type EncodeResponseFunc func(ctx context.Context, w http.ResponseWriter, response interface{}) error
type EncodeErrorFunc func(ctx context.Context, w http.ResponseWriter, err error)

type Handler struct {
	endpoint       endpoint.Endpoint
	decodeRequest  DecodeRequestFunc
	encodeResponse EncodeResponseFunc
	encodeError    EncodeErrorFunc
}

func NewHandler(
	endpoint endpoint.Endpoint,
	decodeRequest DecodeRequestFunc,
	encodeResponse EncodeResponseFunc,
	encodeError EncodeErrorFunc,
) *Handler {
	return &Handler{
		endpoint,
		decodeRequest,
		encodeResponse,
		encodeError,
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req, err := h.decodeRequest(ctx, r)
	if err != nil {
		h.encodeError(ctx, w, err)
		return
	}

	resp, err := h.endpoint(ctx, req)
	if err != nil {
		h.encodeError(ctx, w, err)
		return
	}

	err = h.encodeResponse(ctx, w, resp)
	if err != nil {
		h.encodeError(ctx, w, err)
		return
	}
}
