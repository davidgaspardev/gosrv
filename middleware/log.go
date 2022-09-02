package middleware

import (
	"fmt"
	"gosrv/helpers"

	"github.com/davidgaspardev/golog"
)

func Logger(request *helpers.Request) *MiddlewareFailed {
	address := request.RemoteAddr

	golog.Io("Request", fmt.Sprintf("(%s) Method: %s", address, request.Method))
	golog.Io("Request", fmt.Sprintf("(%s) Path: %s", address, request.URL.Path))

	if len(request.URL.RawQuery) > 0 {
		golog.Io("Request", fmt.Sprintf("(%s) Query: %s", address, request.URL.RawQuery))
	}

	headers := []string{
		"Content-Type",
		"Accept",
		"User-Agent",
		"Authorization",
		"Origin",
		"X-Forwarded-For",
		"X-Real-Ip",
	}

	for _, header := range headers {
		headerValue := request.Header.Get(header)
		if headerValue != "" {
			golog.Io("Request", fmt.Sprintf("(%s) %s: %s", address, header, headerValue))
		}
	}

	return nil
}
