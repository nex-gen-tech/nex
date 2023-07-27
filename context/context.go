package context

import (
	"net/http"
	"sync"

	"github.com/nex-gen-tech/nex/pkg/nexval"
)

// Context is the main structure that will be passed to handlers.
// It provides methods and fields to interact with the HTTP request and response.
type Context struct {
	Request    *http.Request
	Response   http.ResponseWriter
	Params     map[string]string // for route parameters
	Res        *NexResponse
	PathParam  *PathParam
	QueryParam *QueryParam
	Form       *Form
	mu         sync.RWMutex // Mutex for concurrent access to the context fields
}

// NewContext creates a new instance of Context.
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	ctx := Context{
		Request:  r,
		Response: w,
		Params:   make(map[string]string),
	}

	// Response is wrapped in a NexResponse
	ctx.Res = NewNexResponse(&ctx)
	ctx.PathParam = NewPathParam(&ctx)
	ctx.QueryParam = NewQueryParam(&ctx)
	ctx.Form = NewForm(&ctx)

	return &ctx
}

// SetHeader sets a header for the response.
func (c *Context) SetHeader(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Response.Header().Set(key, value)
}

// GetHeader retrieves a header from the request.
func (c *Context) GetHeader(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.Request.Header.Get(key)
}

// Validate - validates a struct.with the Go Playground validator v10 package struct tags.
func (c *Context) Validate(v interface{}) []nexval.ValidationError {
	return nexval.New().Validate(v)
}
