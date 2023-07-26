package context

import (
	"encoding/json"
	"net/http"
	"sync"
)

// Context is the main structure that will be passed to handlers.
// It provides methods and fields to interact with the HTTP request and response.
type Context struct {
	Request  *http.Request
	Response http.ResponseWriter

	// Additional fields can be added as needed, for example:
	Params map[string]string // for route parameters
	// Errors []error           // to collect errors during request processing

	mu sync.RWMutex // Mutex for concurrent access to the context fields
}

// NewContext creates a new instance of Context.
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request:  r,
		Response: w,
		Params:   make(map[string]string),
	}
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

// Additional methods can be added as needed, for example:
// - FormValue(key string) string
// - QueryParam(key string) string
// - SetCookie(cookie *http.Cookie)
// - GetCookie(name string) (*http.Cookie, error)
// - Redirect(status int, url string)
// ... and so on.
