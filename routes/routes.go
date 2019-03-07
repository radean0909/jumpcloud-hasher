package routes

import (
	"net/http"

	"github.com/radean0909/jumpcloud-hasher/middleware"
)

type Router struct {
	Mux    *http.ServeMux
	Routes map[string]func(w http.ResponseWriter, r *http.Request) // Note setting this to a map prevents overloading of routes, that functionality is handled in the controller
}

func NewRouter(mux *http.ServeMux) *Router {
	return &Router{
		Mux:    mux,
		Routes: make(map[string]func(w http.ResponseWriter, r *http.Request)),
	}
}

func (r *Router) AddRoute(path string, process func(w http.ResponseWriter, r *http.Request)) {
	r.Routes[path] = process
}

func (r *Router) ParseRoutes() {
	mw := middleware.ChainMiddleware(middleware.HttpLogger, middleware.HttpTracer)

	for path, process := range r.Routes {
		r.Mux.HandleFunc(path, mw(process))
	}
}
