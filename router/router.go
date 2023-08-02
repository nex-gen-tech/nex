package router

import (
	"context"
	"fmt"
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

	// for route printing
	printRoutes      bool
	printMiddlewares bool
}

// NewRouter returns a new router instance.
func NewRouter() *Router {
	return &Router{
		tree:             NewTree(),
		log:              nexlog.New("NEX-LOG"),
		printRoutes:      false,
		printMiddlewares: false,
	}
}

// SetErrorHandler sets a custom error handler for the router.
func (r *Router) SetErrorHandler(handler func(*nexctx.Context, error)) {
	r.errorHandler = handler
}

// SetPrintRoutes sets the router to print the registered routes on startup.
func (r *Router) SetPrintRoutes(print bool) {
	r.printRoutes = print
}

// SetPrintMiddlewares sets the router to print the registered middlewares on startup.
func (r *Router) SetPrintMiddlewares(print bool) {
	// if printRoutes is false show error message
	if !r.printRoutes {
		r.log.Error("please set SetPrintRoutes to true before SetPrintMiddlewares")
	}

	r.printMiddlewares = print
}

// Group creates a new router group with the specified prefix.
func (r *Router) Group(prefix string) *RouterGroup {
	return &RouterGroup{
		prefix:      prefix,
		router:      r,
		middlewares: make([]MiddlewareFunc, 0),
	}
}

// Use appends middleware(s) to the router's middleware stack.
func (r *Router) Use(middleware ...MiddlewareFunc) {
	r.middlewares = append(r.middlewares, middleware...)
}

func (r *Router) AddRoute(method, path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	r.tree.AddRoute(method, method+":"+path, handler, middlewares...)
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

	allMiddlewares := append(r.middlewares, nodeMiddlewares...)
	// for _, middleware := range allMiddlewares {
	// 	name := runtime.FuncForPC(reflect.ValueOf(middleware).Pointer()).Name()
	// 	log.Printf("middleware: %s", name)
	// 	handler = middleware(handler)
	// }
	// reverse the middleware stack
	for i := len(allMiddlewares) - 1; i >= 0; i-- {
		name := runtime.FuncForPC(reflect.ValueOf(allMiddlewares[i]).Pointer()).Name()
		log.Printf("middleware: %s", name)
		handler = allMiddlewares[i](handler)
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
func (r *Router) Run(addr string) error {
	// Print the list of registered routes
	r.tree.root.printRoutes(printRoutesConfig{Prefix: ""})

	r.Address = addr

	// Define the server
	s := &http.Server{
		Addr:    r.Address,
		Handler: r,
	}

	errCh := make(chan error, 1)

	// Start the server in a goroutine so that it doesn't block
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- fmt.Errorf("error starting server: %s", err.Error())
		}
		close(errCh)
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
		// The context is used to inform the server it has 5 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			return fmt.Errorf("server forced to shutdown: %s", err.Error())
		}

		return nil

	case err := <-errCh:
		return err
	}
}
