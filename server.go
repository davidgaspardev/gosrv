package gosrv

import (
	"fmt"
	"net/http"

	"github.com/davidgaspardev/golog"
	"github.com/davidgaspardev/gosrv/controller"
	"github.com/davidgaspardev/gosrv/middleware"
)

// Alias for the types of the package
type Server = *_Server
type Middleware = middleware.Middleware
type Controller = controller.Controller

// NewServer creates a new server instance
func NewServer() Server {
	mux := http.NewServeMux()

	return &_Server{
		port:   8080,
		mux:    mux,
		router: createRouter(mux),
		logger: false,
	}
}

type _Server struct {
	port   uint16
	mux    *http.ServeMux
	router *_Router
	logger bool
}

func (srv Server) SetPort(port uint16) {
	srv.port = port
}

func (srv Server) SetLogger(show bool) {
	srv.logger = show
}

func (srv Server) Run() error {
	// Build the routes
	srv.router.Build()

	portFormat := fmt.Sprintf(":%d", srv.port)

	if srv.logger {
		golog.System("Server", "Routes builded")
		golog.System("Server", "Listening at "+portFormat)
	}

	return http.ListenAndServe(portFormat, srv.mux)
}

func (srv Server) AddRoute(method, path string, middlewares []Middleware, controller Controller) {
	srv.router.Add(method, path, middlewares, controller)

	if srv.logger {
		golog.System("Server", fmt.Sprintf("Route created: %s (%s)", path, method))
	}
}
