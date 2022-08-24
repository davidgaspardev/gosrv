package helpers

import "net/http"

const (
	// HTTP METHODS
	METHOD_OPTIONS = http.MethodOptions
	METHOD_GET     = http.MethodGet
	METHOD_POST    = http.MethodPost

	STATUS_OK                    = http.StatusOK
	STATUS_NO_CONTENT            = http.StatusNoContent
	STATUS_BAD_REQUEST           = http.StatusBadRequest
	STATUS_INTERNAL_SERVER_ERROR = http.StatusInternalServerError

	// HTTP HEADER
	HEADER_ACCESS_CONTROL_ALLOW_ORIGIN      = "Access-Control-Allow-Origin"
	HEADER_ACCESS_CONTROL_ALLOW_CREDENTIALS = "Access-Control-Allow-Credentials"
	HEADER_ACCESS_CONTROL_ALLOW_METHODS     = "Access-Control-Allow-Methods"
	HEADER_ACCESS_CONTROL_ALLOW_HEADERS     = "Access-Control-Allow-Headers"
	HEADER_ACCEPT                           = "Accept"
	HEADER_AUTHORIZATION                    = "Authorization"
	HEADER_CONTENT_TYPE                     = "Content-Type"
	HEADER_ORIGIN                           = "Origin"
	HEADER_APPLICATION_JSON                 = "application/json"
)
