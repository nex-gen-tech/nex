package nex

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"path"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/fatih/color"
)

type (
	HandlerFunc    func(*Context)
	MiddlewareFunc func(HandlerFunc) HandlerFunc
)

// Route represents a single route in the router.
type Route struct {
	Method     string
	Pattern    *regexp.Regexp
	Handler    HandlerFunc
	Params     []string
	OrigPath   string
	Middleware []MiddlewareFunc
}

// Router is a HTTP request router.
type Router struct {
	Routes     []*Route
	Middleware []MiddlewareFunc
	Address    string
}

// New creates a new Router.
func NewRouter() *Router {
	return &Router{}
}

// New creates a new Router.
func New() *Router {
	return &Router{}
}

// Run starts the server with the given address.
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

	log.Println("Server exiting")
}

func (r *Router) AddRoute(
	method string,
	pattern string,
	handler HandlerFunc,
	middleware ...MiddlewareFunc,
) {
	regex, params := compilePattern(pattern)

	for _, route := range r.Routes {
		if route.Method == method && route.Pattern.String() == regex.String() {
			panic(fmt.Sprintf("Duplicate route: %s %s", method, pattern))
		}
	}

	route := &Route{
		Pattern:    regex,
		Params:     params,
		Method:     method,
		Handler:    handler,
		OrigPath:   pattern,
		Middleware: middleware,
	}
	r.Routes = append(r.Routes, route)
}

// GET registers a new GET route.
func (r *Router) GET(pattern string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	r.AddRoute("GET", pattern, handler, middleware...)
}

// POST registers a new POST route.
func (r *Router) POST(pattern string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	r.AddRoute("POST", pattern, handler, middleware...)
}

// PATCH - registers a new PATCH route.
func (r *Router) PATCH(pattern string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	r.AddRoute("PATCH", pattern, handler, middleware...)
}

// DELETE - registers a new DELETE route.
func (r *Router) DELETE(pattern string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	r.AddRoute("DELETE", pattern, handler, middleware...)
}

// PUT - registers a new PUT route.
func (r *Router) PUT(pattern string, handler HandlerFunc) {
	r.AddRoute("PUT", pattern, handler)
}

// compilePattern - compiles a route pattern into a regular expression.
func compilePattern(pattern string) (*regexp.Regexp, []string) {
	segments := strings.Split(pattern, "/")
	params := make([]string, 0)

	for i, segment := range segments {
		if strings.HasPrefix(segment, ":") {
			params = append(params, segment[1:])
			segments[i] = "([^/]+)"
		}
	}

	pattern = "^" + strings.Join(segments, "/") + "$" // Add ^ at the beginning and $ at the end
	regex, regexErr := regexp.Compile(pattern)
	if regexErr != nil {
		panic(regexErr)
	}

	return regex, params
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateRequestID() string {
	return fmt.Sprintf("%08d", rand.Intn(100000000)) // generate a random 8-digit number
}

// ServeHTTP matches the request to a route and calls the corresponding handler.
// Find and execute the route matching the request.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.Routes {
		if matches := route.Pattern.FindStringSubmatch(req.URL.Path); matches != nil {
			if req.Method != route.Method {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			params := make(map[string]string)
			for i, match := range matches[1:] {
				params[route.Params[i]] = match
			}

			// Create a new statusWriter
			sw := ResponseWriteNex{ResponseWriter: w}

			c := NewContext(req, sw, params)

			// Generate and set request ID if not present
			reqID := req.Header.Get("X-Request-ID")
			if reqID == "" {
				reqID = generateRequestID()
				req.Header.Set("X-Request-ID", reqID)
			}
			c.Set("requestID", reqID)

			// Combine router-level and route-level middleware
			allMiddleware := append(r.Middleware, route.Middleware...)

			// Wrap the route handler with the middleware
			h := route.Handler
			for _, mw := range allMiddleware {
				h = mw(h)
			}

			// Call the handler
			h(c)

			// After the handler has run, the status code should have been set.
			// We can store it in the context.
			status := sw.status
			if status == 0 {
				status = 200 // Default status code is 200 OK
			}
			c.Set("status", status)

			return
		}
	}

	http.NotFound(w, req)
}

// Use - adds a middleware to the router
func (r *Router) Use(middleware ...MiddlewareFunc) {
	r.Middleware = append(r.Middleware, middleware...)
}

// Static - serves static files from the given directory
func (r *Router) Static(pattern string, dir string) {
	pattern = path.Join(pattern, "/")
	fs := http.FileServer(http.Dir(dir))
	r.GET(pattern+"*", func(ctx *Context) {
		fs.ServeHTTP(ctx.Res, ctx.Req)
	})
}

// PrintRoutes - Prints the Routes
// if file is true, it prints the routes to a file called routes.json in root of the project
func (r *Router) PrintRoutes(file ...bool) {
	fmt.Println()
	fmt.Println("Printing routes:")
	fmt.Println()
	routes := make([]Route, 0)
	for i, route := range r.Routes {
		method := color.New(color.FgHiMagenta).Sprintf(" (%s) ", route.Method)
		index := color.New(color.FgHiMagenta).Sprint(i)
		fmt.Printf("%s -> %s %s\n", index, route.OrigPath, method)
		if len(file) > 0 && file[0] {
			routes = append(routes, *route)
		}
	}
	r.writeRouteToFile(&routes)
	fmt.Println()
}

// writeRouteToFile - writes the route to a file
func (r *Router) writeRouteToFile(route *[]Route) {
	_, err := os.Stat("routes.csv")
	if err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			_, err := os.Create("routes.csv")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}

	// Write to file and don't append
	f, err := os.OpenFile("routes.csv", os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}

	// Write to file
	for _, route := range *route {
		_, err := f.WriteString(route.OrigPath + "," + route.Method + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	defer f.Close()
}

//

// RouterGroup represents a group of routes with a common prefix.
type RouterGroup struct {
	router     *Router
	prefix     string
	Middleware []MiddlewareFunc
}

// NewRouterGroup creates a new RouterGroup with the given prefix.
func (r *Router) Group(prefix string) *RouterGroup {
	return &RouterGroup{
		router: r,
		prefix: prefix,
	}
}

// AddRoute adds a new route to the router group.
func (rg *RouterGroup) AddRoute(method string, pattern string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	// Ensure prefix and pattern start with "/"
	if !strings.HasPrefix(rg.prefix, "/") {
		rg.prefix = "/" + rg.prefix
	}
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}

	// Combine group-level and route-level middleware
	allMiddleware := append(rg.Middleware, middleware...)

	rg.router.AddRoute(method, rg.prefix+pattern, handler, allMiddleware...)
}

// GET registers a new GET route in the router group.
func (rg *RouterGroup) GET(pattern string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	rg.AddRoute("GET", pattern, handler, middleware...)
}

// POST registers a new POST route in the router group.
func (rg *RouterGroup) POST(pattern string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	rg.AddRoute("POST", pattern, handler, middleware...)
}

// PATCH registers a new PATCH route in the router group.
func (rg *RouterGroup) PATCH(pattern string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	rg.AddRoute("PATCH", pattern, handler, middleware...)
}

// DELETE registers a new DELETE route in the router group.
func (rg *RouterGroup) DELETE(pattern string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	rg.AddRoute("DELETE", pattern, handler, middleware...)
}

// PUT registers a new PUT route in the router group.
func (rg *RouterGroup) PUT(pattern string, handler HandlerFunc, middleware ...MiddlewareFunc) {
	rg.AddRoute("PUT", pattern, handler, middleware...)
}

// PrintRoutes - prints all routes in the router group.
func (rg *RouterGroup) PrintRoutes(file ...bool) {
	rg.router.PrintRoutes(file...)
}

// Group creates a new RouterGroup with the given prefix.
func (rg *RouterGroup) Group(prefix string) *RouterGroup {
	newPrefix := rg.prefix
	if !strings.HasSuffix(newPrefix, "/") {
		newPrefix += "/"
	}
	return &RouterGroup{
		router: rg.router,
		prefix: newPrefix + prefix,
	}
}
