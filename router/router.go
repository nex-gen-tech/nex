package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	nexctx "github.com/nex-gen-tech/nex/context"
)

const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
	MethodPatch  = "PATCH"
)

type HandlerFunc func(*nexctx.Context)

type Router struct {
	tree    *Tree
	Address string
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
	ctx := nexctx.NewContext(w, req)
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
			log.Fatalf("listen: %s\n", err)
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
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Existing Server...")
}
