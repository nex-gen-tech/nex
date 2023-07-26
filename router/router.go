package router

import (
	"net/http"

	"github.com/nex-gen-tech/nex/context"
)

const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
	MethodPatch  = "PATCH"
)

type HandlerFunc func(*context.Context)

type Router struct {
	tree *Tree
}

func NewRouter() *Router {
	return &Router{
		tree: NewTree(),
	}
}

func (r *Router) AddRoute(method, path string, handler HandlerFunc) {
	r.tree.AddRoute(method+":"+path, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, params := r.tree.Match(req.Method + ":" + req.URL.Path)
	if handler == nil {
		http.NotFound(w, req)
		return
	}
	ctx := context.NewContext(w, req)
	for key, value := range params {
		ctx.Params[key] = value
	}
	handler(ctx)
}

func (r *Router) GET(path string, handler HandlerFunc) {
	r.AddRoute(MethodGet, path, handler)
}

func (r *Router) POST(path string, handler HandlerFunc) {
	r.AddRoute(MethodPost, path, handler)
}

func (r *Router) PUT(path string, handler HandlerFunc) {
	r.AddRoute(MethodPut, path, handler)
}

func (r *Router) DELETE(path string, handler HandlerFunc) {
	r.AddRoute(MethodDelete, path, handler)
}

func (r *Router) PATCH(path string, handler HandlerFunc) {
	r.AddRoute(MethodPatch, path, handler)
}

func (r *Router) Run(addr string) error {
	return http.ListenAndServe(addr, r)
}
