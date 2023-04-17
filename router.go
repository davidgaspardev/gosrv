package gosrv

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/davidgaspardev/gosrv/helpers"
)

type _RouteConfig struct {
	paramsPosition []uint8
	middlewares    []Middleware
	controller     Controller
}

type _Router struct {
	routes map[ /** path */ string]map[ /** method */ string]*_RouteConfig
	mux    *http.ServeMux
}

func createRouter(mux *http.ServeMux) *_Router {
	return &_Router{
		mux:    mux,
		routes: make(map[string]map[string]*_RouteConfig),
	}
}

func (router *_Router) Add(method string, path string, middlewares []Middleware, controller Controller) {
	route := _RouteConfig{
		paramsPosition: nil,
		middlewares:    middlewares,
		controller:     controller,
	}

	paramsPosition := make([]uint8, strings.Count(path, "/:"))

	if len(paramsPosition) > 0 {
		dirs := strings.Split(path, "/")
		paramIndex := uint8(0)
		index := uint8(0)
		length := uint8(len(dirs))

		for ; index < length; index++ {
			if strings.HasPrefix(dirs[index], ":") {
				paramsPosition[paramIndex] = index
				paramIndex++
			}
		}

		route.paramsPosition = paramsPosition
	}

	// // Get index from variables in the path (URL)
	// var dynamicPosition []uint8
	// dirs := strings.Split(path, "/")
	// for i := uint8(0); i < uint8(len(dirs)); i++ {
	// 	if strings.HasPrefix(dirs[i], ":") {
	// 		dynamicPosition = append(dynamicPosition, i)
	// 	}
	// }

	// if len(dynamicPosition) > 0 {
	// 	params := make([]*_Param, len(dynamicPosition))
	// 	for i := 0; i < len(params); i++ {
	// 		params[]
	// 	}
	// 	route.dynamicPath = dynamicPath
	// }

	if router.routes[path] == nil {
		router.routes[path] = make(map[string]*_RouteConfig)
	}

	router.routes[path][method] = &route
}

func (router *_Router) Build() {
	router.mux.HandleFunc("/", func(responseWriter http.ResponseWriter, requestCode *http.Request) {
		response := &helpers.Response{ResponseWriter: responseWriter}
		request := &helpers.Request{Request: requestCode}

		if request.HasRequestOrigin() {
			origin := request.GetOrigin()
			response.AddCors(origin)

			if request.Method == http.MethodOptions {
				response.NoContent()
				return
			}
		}

		// Get route
		route := router.routes[request.URL.Path]
		var routeConfig *_RouteConfig
		if route == nil {
			// Dynamic path
			routeConfig = router.findDynamicPathRouteByRequest(request)
			if routeConfig == nil {
				response.NotFound()
				return
			}
		} else {
			// Non-dynamic path
			routeConfig = route[request.Method]
			if routeConfig == nil {
				response.BadRequest(fmt.Errorf("http method invalid"))
				return
			}
		}

		// Run middleware
		for i := 0; i < len(routeConfig.middlewares); i++ {
			if middlewareFailed := routeConfig.middlewares[i](request); middlewareFailed != nil {
				switch middlewareFailed.Code {
				case http.StatusBadRequest:
					response.BadRequest(middlewareFailed.Error)
				case http.StatusUnauthorized:
					response.Unauthorized(middlewareFailed.Error)
				case http.StatusForbidden:
					response.Forbidden(middlewareFailed.Error)
				default:
					response.InternalServerError(middlewareFailed.Error)
				}
				return
			}
		}

		// Run controller
		routeConfig.controller(request, response)
	})
}

func (router *_Router) findDynamicPathRouteByRequest(request *helpers.Request) *_RouteConfig {
	requestPathDirs := strings.Split(request.URL.Path, "/")

	for path, route := range router.routes {
		routeConfig := route[request.Method]
		if routeConfig == nil || len(routeConfig.paramsPosition) < 1 {
			continue
		}

		pathDirs := strings.Split(path, "/")
		if len(requestPathDirs) != len(pathDirs) {
			continue
		}

		found := true
		for position := uint8(0); position < uint8(len(requestPathDirs)); position++ {
			if !contains(routeConfig.paramsPosition, position) {
				if requestPathDirs[position] != pathDirs[position] {
					found = false
					break
				}
			} else {
				request.AddParam(pathDirs[position][1:], requestPathDirs[position])
			}
		}

		if found {
			return routeConfig
		}
	}

	return nil
}

func contains(s []uint8, e uint8) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
