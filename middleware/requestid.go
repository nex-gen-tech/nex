package middleware

import (
	"context"

	"github.com/google/uuid"
	nexctx "github.com/nex-gen-tech/nex/context" // Assuming the context is now in this package
	"github.com/nex-gen-tech/nex/router"         // Assuming the router is in this package
)

const (
	requestIDKey    = "RequestID"
	requestIDHeader = "X-Request-ID"
)

// RequestID is a middleware that assigns a unique ID to each incoming request.
func RequestID() router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *nexctx.Context) {
			// Check if request header already has a Request ID
			reqID := c.Request.Header.Get(requestIDHeader)
			if reqID == "" {
				// Generate a new UUID for the request
				reqID = uuid.New().String()
			}

			// Add the Request ID to the request's context
			ctx := context.WithValue(c.Request.Context(), requestIDKey, reqID)
			c.Request = c.Request.WithContext(ctx)

			// Optionally, set the Request ID in the response header
			c.Response.Header().Set(requestIDHeader, reqID)

			// Continue processing the request
			next(c)
		}
	}
}

// GetRequestID retrieves the Request ID from the context.
func GetRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		return reqID
	}
	return ""
}
