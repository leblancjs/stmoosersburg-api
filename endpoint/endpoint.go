package endpoint

import "context"

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

type Middleware func(endpoint Endpoint) Endpoint
