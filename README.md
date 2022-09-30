# Go SRV

Go server or `Go SRV` is a http server minimal.

## Using

```go
package main

import "github.com/davidgaspardev/gosrv.go"

func main() {
    // Create HTTP server
    server := gosrv.NewServer()

    server.AddRoute("GET", "/api/v1/hello/:name", []middleware.Middleware{
        // Print incoming HTTP requests
        middleware.Logger,
    }, HelloWorld)

    server.Run()
}

// Hello World Controller
func HelloWorld(req *helpers.Request, res *helpers.Response) {
    // Get path parameters from request
	name := req.GetParam("name")

	res.Ok("Hello World for you, "+name)
}
```

## Files tree

Below is the tree diagram of the project files.

```bash
 ---gosrv
    |   go.mod
    |   go.sum
    |   gosrv.go
    |   LICENSE
    |   Makefile
    |   README.md
    |   router.go
    |   server.go
    |
    +---controller
    |       controller.go
    |
    +---examples
    |       main.go
    |
    +---helpers
    |       constants.go
    |       request.go
    |       response.go
    |
    +---middleware
    |       cors.go
    |       log.go
    |       middleware.go
    |       validation.go
    |
    +---model
    |       data.go
    |
    \---tests
            multi_requests.py
```