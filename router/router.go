package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"syscall"
	"time"

	nexctx "github.com/nex-gen-tech/nex/context"
	"github.com/nex-gen-tech/nexlog"
)

const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
	MethodPatch  = "PATCH"
)

// HandlerFunc defines a function to serve HTTP requests.
type HandlerFunc func(*nexctx.Context)

// Router is a http.Handler which can be used to dispatch requests to different
type Router struct {
	tree         *Tree
	Address      string
	log          nexlog.Logger
	errorHandler func(*nexctx.Context, error)
	middlewares  []MiddlewareFunc
}

// NewRouter returns a new router instance.
func NewRouter() *Router {
	return &Router{
		tree: NewTree(),
		log:  nexlog.New("NEX-LOG"),
	}
}

// SetErrorHandler sets a custom error handler for the router.
func (r *Router) SetErrorHandler(handler func(*nexctx.Context, error)) {
	r.errorHandler = handler
}

// Use appends middleware(s) to the router's middleware stack.
func (r *Router) Use(middleware ...MiddlewareFunc) {
	r.middlewares = append(r.middlewares, middleware...)
}

func (r *Router) AddRoute(method, path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	r.tree.AddRoute(method+":"+path, handler, middlewares...)
}

// ServeHTTP implements the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := nexctx.NewContext(w, req)
	handler, params, nodeMiddlewares := r.tree.Match(req.Method + ":" + req.URL.Path)

	if handler == nil {
		http.NotFound(w, req)
		return
	}

	for key, value := range params {
		ctx.Params[key] = value
	}

	allMiddlewares := append(nodeMiddlewares, r.middlewares...)
	for _, middleware := range allMiddlewares {
		name := runtime.FuncForPC(reflect.ValueOf(middleware).Pointer()).Name()
		log.Printf("middleware: %s", name)
		handler = middleware(handler)
	}

	// Execute the handler
	handler(ctx)

	// Check if an error was set in the context
	if ctx.Error() != nil && r.errorHandler != nil {
		r.errorHandler(ctx, ctx.Error())
	}
}

// GET is a shortcut for router.AddRoute("GET", path, handler)
func (r *Router) GET(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	r.AddRoute(MethodGet, path, handler, middlewares...)
}

// POST is a shortcut for router.AddRoute("POST", path, handler)
func (r *Router) POST(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	r.AddRoute(MethodPost, path, handler, middlewares...)
}

// PUT is a shortcut for router.AddRoute("PUT", path, handler)
func (r *Router) PUT(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	r.AddRoute(MethodPut, path, handler, middlewares...)
}

// DELETE is a shortcut for router.AddRoute("DELETE", path, handler)
func (r *Router) DELETE(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	r.AddRoute(MethodDelete, path, handler, middlewares...)
}

// PATCH is a shortcut for router.AddRoute("PATCH", path, handler)
func (r *Router) PATCH(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	r.AddRoute(MethodPatch, path, handler, middlewares...)
}

// Run starts the HTTP server.
func (r *Router) Run(addr string) {
	r.Address = addr

	// Define the server
	s := &http.Server{
		Addr:    r.Address,
		Handler: r,
	}

	// Start the server in a goroutine so that it doesn't block
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			r.log.FatalF("error starting server: %s", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	r.log.InfoF("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		r.log.FatalF("Server forced to shutdown: %s", err.Error())
	}

	r.log.InfoF("Existing Server...")
}
