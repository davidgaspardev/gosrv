package middleware

import "github.com/davidgaspardev/gosrv/helpers"

/// Middleware failed struture is used to build http response
type MiddlewareFailed struct {
	Code  uint16 // status code (http)
	Error error
}

type Middleware = func(request *helpers.Request) *MiddlewareFailed
