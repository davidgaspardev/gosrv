package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/davidgaspardev/gosrv/helpers"
)

func SimpleCorsMiddleware(req *helpers.Request) *MiddlewareResponse {
	if req.GetOrigin() == "" {
		return &MiddlewareResponse{
			Code:  http.StatusForbidden,
			Error: errors.New("origin not allowed"),
		}
	}

	responseHeader := getResponseHeader(req.Context())
	responseHeader.Set(helpers.HEADER_ACCESS_CONTROL_ALLOW_ORIGIN, req.GetOrigin())

	return nil
}

func getResponseHeader(ctx context.Context) http.Header {
	value, ok := helpers.GetFromContext(ctx, helpers.ResponseHeaderKey)
	if ok {
		header, ok := value.(http.Header)
		if !ok {
			return nil
		}

		return header
	}

	return make(http.Header)
}
