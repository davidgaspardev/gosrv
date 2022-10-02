package gosrv

import (
	"fmt"
	"net/http"

	"github.com/davidgaspardev/golog"
)

func NewServer() Server {
	mux := http.NewServeMux()
	return &_Server{
		port:   8080,
		mux:    mux,
		router: createRouter(mux),
	}
}

type _Server struct {
	port   uint16
	mux    *http.ServeMux
	router *_Router
}

func (srv *_Server) SetPort(port uint16) {
	srv.port = port
}

func (srv *_Server) SetLogger(show bool) {
	logger = show
}

func (srv *_Server) Run() error {
	// Build the routes
	srv.router.Build()

	portFormat := fmt.Sprintf(":%d", srv.port)

	if logger {
		golog.System("Server", "Routes builded")
		golog.System("Server", "Listening at "+portFormat)
	}

	return http.ListenAndServe(portFormat, srv.mux)
}

func (srv *_Server) AddRoute(method, path string, middlewares []Middleware, controller Controller) {
	srv.router.Add(method, path, middlewares, controller)

	if logger {
		golog.System("Server", fmt.Sprintf("Route created: %s (%s)", path, method))
	}
}
