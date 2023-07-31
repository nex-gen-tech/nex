package router

type MiddlewareFunc func(HandlerFunc) HandlerFunc

// RouterGroup allows grouping of routes with shared path prefix and middleware.
type RouterGroup struct {
	prefix      string
	router      *Router
	middlewares []MiddlewareFunc
}

// NewGroup initializes a new router group with the provided prefix.
func (r *Router) Group(prefix string) *RouterGroup {
	return &RouterGroup{
		prefix: prefix,
		router: r,
	}
}

// Use appends middleware(s) to the group's middleware stack.
// These middlewares will be executed in the order they are added for all routes in this group.
func (group *RouterGroup) Use(middleware ...MiddlewareFunc) {
	group.middlewares = append(group.middlewares, middleware...)
}

// addRoute is an internal method to handle adding routes with the provided method, path, handler, and middlewares.
func (group *RouterGroup) addRoute(method, path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	allMiddlewares := append(group.middlewares, middlewares...)
	group.router.AddRoute(method, group.prefix+path, handler, allMiddlewares...)
}

// GET adds a new route with the GET method to the group.
func (group *RouterGroup) GET(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	group.addRoute(MethodGet, path, handler, middlewares...)
}

// POST adds a new route with the POST method to the group.
func (group *RouterGroup) POST(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	group.addRoute(MethodPost, path, handler, middlewares...)
}

// PUT adds a new route with the PUT method to the group.
func (group *RouterGroup) PUT(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	group.addRoute(MethodPut, path, handler, middlewares...)
}

// DELETE adds a new route with the DELETE method to the group.
func (group *RouterGroup) DELETE(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	group.addRoute(MethodDelete, path, handler, middlewares...)
}

// PATCH adds a new route with the PATCH method to the group.
func (group *RouterGroup) PATCH(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	group.addRoute(MethodPatch, path, handler, middlewares...)
}
