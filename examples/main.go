package main

import (
	"gosrv"
	"gosrv/helpers"
	"gosrv/middleware"
)

func main() {
	server := gosrv.NewServer()

	server.SetPort(8080)
	server.SetLogger(true)

	server.AddRoute("GET", "/v1/hello", []middleware.Middleware{}, HelloWorld)
	server.AddRoute("GET", "/v1/hello/:name", nil, HelloWithParam)

	server.AddRoute("POST", "/v1/world", []middleware.Middleware{
		middleware.HasQuery("name"),
	}, WorldWithQuery)

	server.Run()
}

func HelloWorld(req *helpers.Request, res *helpers.Response) {
	res.Ok("hello world")
}

func HelloWithParam(req *helpers.Request, res *helpers.Response) {
	data := map[string]interface{}{
		"name": req.GetParam("name"),
		"type": req.GetContentType(),
	}

	res.Ok(data)
}

func WorldWithQuery(req *helpers.Request, res *helpers.Response) {
	res.NoContent()
}
