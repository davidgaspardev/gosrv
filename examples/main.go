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

	server.Run()
}

func HelloWorld(req *helpers.Request, res *helpers.Response) {
	res.Ok("hello world")
}
