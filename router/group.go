package router

type MiddlewareFunc func(HandlerFunc) HandlerFunc

// RouterGroup allows grouping of routes with shared path prefix and middleware.
type RouterGroup struct {
	prefix      string
	router      *Router
	middlewares []MiddlewareFunc
}

// NewGroup initializes a new router group with the provided prefix.
func (r *Router) NewGroup(prefix string) *RouterGroup {
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

// GET adds a new route with the GET method to the group.
func (group *RouterGroup) GET(path string, handler HandlerFunc) {
	group.router.AddRoute(MethodGet, group.prefix+path, group.applyMiddlewares(handler))
}

// POST adds a new route with the POST method to the group.
func (group *RouterGroup) POST(path string, handler HandlerFunc) {
	group.router.AddRoute(MethodPost, group.prefix+path, group.applyMiddlewares(handler))
}

// PUT adds a new route with the PUT method to the group.
func (group *RouterGroup) PUT(path string, handler HandlerFunc) {
	group.router.AddRoute(MethodPut, group.prefix+path, group.applyMiddlewares(handler))
}

// DELETE adds a new route with the DELETE method to the group.
func (group *RouterGroup) DELETE(path string, handler HandlerFunc) {
	group.router.AddRoute(MethodDelete, group.prefix+path, group.applyMiddlewares(handler))
}

// PATCH adds a new route with the PATCH method to the group.
func (group *RouterGroup) PATCH(path string, handler HandlerFunc) {
	group.router.AddRoute(MethodPatch, group.prefix+path, group.applyMiddlewares(handler))
}

// applyMiddlewares wraps the provided handler with the group's middlewares.
// It applies the middlewares in the order they were added to the group.
func (group *RouterGroup) applyMiddlewares(handler HandlerFunc) HandlerFunc {
	finalHandler := handler

	for i := len(group.middlewares) - 1; i >= 0; i-- {
		finalHandler = group.middlewares[i](finalHandler)
	}

	return finalHandler
}
