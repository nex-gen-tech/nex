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
	Body       *Body
	mu         sync.RWMutex // Mutex for concurrent access to the context fields
	Data       map[string]any
	err        error
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
	ctx.Body = NewBody(&ctx)

	return &ctx
}

// Error returns the error set in the context.
func (c *Context) Error() error {
	return c.err
}

// SetError sets an error in the context.
func (c *Context) SetError(err error) {
	c.err = err
}

// String writes a string response to the client.
func (c *Context) String(status int, s string) {
	c.Response.WriteHeader(status)
	c.Response.Write([]byte(s))
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

// Set -Set the data in context
func (c *Context) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.Data == nil {
		c.Data = make(map[string]any)
	}
	c.Data[key] = value
}

// Get - Get the data from context
func (c *Context) Get(key string) any {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.Data == nil {
		return nil
	}
	return c.Data[key]
}
