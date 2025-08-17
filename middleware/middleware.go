package middleware

import "github.com/davidgaspardev/gosrv/helpers"

// MiddlewareResponse represents a short-circuit response returned by middleware.
//
// If a middleware returns a non-nil *MiddlewareResponse, it indicates that the
// request flow should be interrupted and an HTTP response should be sent immediately.
//
// Convention:
//   - Return nil → continue request flow.
//   - Return *MiddlewareResponse → stop flow and write this response.
//
// Fields:
//
//	Code  - HTTP status code to send.
//	Error - Optional error information (can be nil).
//	Data  - Optional payload (can be serialized to JSON, etc).
type MiddlewareResponse struct {
	Code  uint16 // Status code (http)
	Error error  // Optional
	Data  any    // Optional, payload to response
}

// Middleware defines the signature of a middleware function.
// It can either continue the flow (by returning nil) or short-circuit the flow
// by returning a *MiddlewareResponse.
type Middleware = func(request *helpers.Request) *MiddlewareResponse
