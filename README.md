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