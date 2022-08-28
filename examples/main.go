package main

import (
	"fmt"

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
	server.AddRoute("POST", "/v1/world/:country_name", []middleware.Middleware{
		middleware.Or(
			middleware.HasQuery("language"),
			middleware.HasQuery("president"),
		),
		middleware.HasQuery("currency"),
	}, WorldWithQueryAndParam)

	server.Run()
}

func HelloWorld(req *helpers.Request, res *helpers.Response) {
	res.Ok("hello world")
}

func HelloWithParam(req *helpers.Request, res *helpers.Response) {
	// Get path parameters from request
	name := req.GetParam("name")

	res.Ok(fmt.Sprintf("hello %s", name))
}

func WorldWithQuery(req *helpers.Request, res *helpers.Response) {
	res.NoContent()
}

func WorldWithQueryAndParam(req *helpers.Request, res *helpers.Response) {
	if req.GetAccept() != "application/json" && req.GetAccept() != "" {
		res.BadRequest(fmt.Errorf("accept don't supported"))
		return
	}

	// Get path parameters from request
	countryName := req.GetParam("country_name")

	// Get query parameters from request
	language := req.URL.Query().Get("language")
	president := req.URL.Query().Get("president")
	currency := req.URL.Query().Get("currency")

	// Building response data
	data := make(map[string]interface{})
	data["countryName"] = countryName
	data["language"] = language
	data["president"] = president
	data["currency"] = currency

	res.Ok(data)
}
