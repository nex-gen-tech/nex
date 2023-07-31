package router

import (
	"log"
	"strings"
)

// node is a node in the tree structure.
type node struct {
	path        string
	children    []*node
	handler     HandlerFunc
	param       *paramMatcher
	middlewares []MiddlewareFunc
}

// insert inserts a new node in the tree.
func (n *node) insert(segments []string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	if len(segments) == 0 {
		n.handler = handler
		n.middlewares = middlewares
		return
	}

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
	child.insert(segments[1:], handler, middlewares...)
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
func (t *Tree) AddRoute(path string, handler HandlerFunc, middlewares ...MiddlewareFunc) {
	segments := splitPath(path)
	t.root.insert(segments, handler, middlewares...)
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
