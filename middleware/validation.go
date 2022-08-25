package middleware

import (
	"fmt"
	"gosrv/helpers"
)

func HasQuery(queries ...string) Middleware {
	return func(request *helpers.Request) *MiddlewareFailed {
		queryNum := len(queries)

		for i := 0; i < queryNum; i++ {
			query := request.URL.Query().Get(queries[i])

			if query == "" {
				return &MiddlewareFailed{
					Code:  400,
					Error: fmt.Errorf("missing query: %s", queries[i]),
				}
			}
		}

		return nil
	}
}

func Or(middlewares ...Middleware) Middleware {
	return func(request *helpers.Request) (failed *MiddlewareFailed) {
		middlewaresNum := len(middlewares)

		for i := 0; i < middlewaresNum; i++ {
			if failed = middlewares[i](request); failed == nil {
				break
			}
		}

		return
	}
}
