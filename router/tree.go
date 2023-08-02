package router

import (
	"log"
	"reflect"
	"runtime"
	"strings"
)

// node is a node in the tree structure.
type node struct {
	path        string
	children    []*node
	handler     HandlerFunc
	param       *paramMatcher
	middlewares []MiddlewareFunc
	method      string
}

// insert inserts a new node in the tree.
func (n *node) insert(segments []string, method string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	if len(segments) == 0 {
		n.handler = handler
		n.middlewares = middlewares
		return
	}

	n.method = method
	segment := segments[0]
	child := n.matchChild(segment)
	if child == nil {
		child = &node{path: segment}
		var err error
		child.param, err = newParamMatcher(segment)
		if err != nil {
			log.Fatalf("Failed to create param matcher: %v", err)
			return
		}
		n.children = append(n.children, child)
	}
	child.insert(segments[1:], method, handler, middlewares...)
}

// matchChild matches a child node with the given segment.
func (n *node) matchChild(segment string) *node {
	for _, child := range n.children {
		if child.path == segment {
			return child
		}
		if child.param != nil {
			matched, _ := child.param.match(segment)
			if matched {
				return child
			}
		}
	}
	return nil
}

// search searches for a node in the tree.
func (n *node) search(segments []string, params map[string]string) (HandlerFunc, []MiddlewareFunc) {
	if len(segments) == 0 || (len(segments) == 1 && segments[0] == "") {
		return n.handler, n.middlewares
	}

	segment := segments[0]
	for _, child := range n.children {
		if child.path == segment {
			return child.search(segments[1:], params)
		}
		if child.param != nil {
			matched, value := child.param.match(segment)
			if matched {
				params[child.param.name] = value
				return child.search(segments[1:], params)
			}
		}
	}
	return nil, nil
}

// Tree is a tree structure that holds the routes.
type Tree struct {
	root *node
}

// NewTree creates a new tree.
func NewTree() *Tree {
	return &Tree{root: &node{}}
}

// AddRoute adds a new route to the tree.
func (t *Tree) AddRoute(method, path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	segments := splitPath(path)
	t.root.insert(segments, method, handler, middlewares...) // Add method here
}

// Match matches a path against the tree.
func (t *Tree) Match(path string) (HandlerFunc, map[string]string, []MiddlewareFunc) {
	segments := splitPath(path)
	params := make(map[string]string)
	handler, middlewares := t.root.search(segments, params)
	return handler, params, middlewares
}

// splitPath splits a path into segments.
func splitPath(path string) []string {
	return strings.Split(strings.Trim(path, "/"), "/")
}

// printRoutes config
type printRoutesConfig struct {
	// PrintRoutes prints the routes in the tree.
	PrintRoutes      bool
	PrintMiddlewares bool
	Prefix           string
}

// printRoutes prints the routes in the tree in a recursive manner.
func (n *node) printRoutes(config printRoutesConfig) {
	fullPath := config.Prefix + n.path
	if n.handler != nil {
		log.Printf("Route: %s", fullPath)
		if config.PrintMiddlewares {
			for _, middleware := range n.middlewares {
				name := runtime.FuncForPC(reflect.ValueOf(middleware).Pointer()).Name()
				log.Printf("\tMiddleware: %s", name)
			}
		}
	}
	for _, child := range n.children {
		child.printRoutes(printRoutesConfig{
			Prefix: fullPath + "/",
		})
	}
}
