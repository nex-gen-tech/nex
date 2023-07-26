package router

import (
	"log"
	"strings"
)

type node struct {
	path     string
	children []*node
	handler  HandlerFunc
	param    *paramMatcher
}

func (n *node) insert(segments []string, handler HandlerFunc) {
	if len(segments) == 0 {
		n.handler = handler
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
	child.insert(segments[1:], handler)
}

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

func (n *node) search(segments []string, params map[string]string) HandlerFunc {
	if len(segments) == 0 || (len(segments) == 1 && segments[0] == "") {
		return n.handler
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
	return nil
}

type Tree struct {
	root *node
}

func NewTree() *Tree {
	return &Tree{root: &node{}}
}

func (t *Tree) AddRoute(path string, handler HandlerFunc) {
	segments := splitPath(path)
	t.root.insert(segments, handler)
}

func (t *Tree) Match(path string) (HandlerFunc, map[string]string) {
	segments := splitPath(path)
	params := make(map[string]string)
	handler := t.root.search(segments, params)
	return handler, params
}

func splitPath(path string) []string {
	return strings.Split(strings.Trim(path, "/"), "/")
}
