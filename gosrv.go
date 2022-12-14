package gosrv

import (
	"github.com/davidgaspardev/gosrv/controller"
	"github.com/davidgaspardev/gosrv/middleware"
)

type Middleware = middleware.Middleware
type Controller = controller.Controller

type Server interface {
	SetPort(port uint16)
	SetLogger(show bool)

	AddRoute(method string, path string, middlewares []Middleware, controller Controller)

	Run() error
}

var logger bool
