package context

import (
	"encoding/json"
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

// JSON sends a JSON response with the given status code and payload.
func (c *Context) JSON(status int, payload interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Response.Header().Set("Content-Type", "application/json")
	c.Response.WriteHeader(status)
	if err := json.NewEncoder(c.Response).Encode(payload); err != nil {
		http.Error(c.Response, err.Error(), http.StatusInternalServerError)
	}
}

// Text sends a plain text response with the given status code and payload.
func (c *Context) Text(status int, payload string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Response.Header().Set("Content-Type", "text/plain")
	c.Response.WriteHeader(status)
	c.Response.Write([]byte(payload))
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
