package nexmiddleware

import (
	"net/http"

	"github.com/nex-gen-tech/nex/context" // Assuming the context is now in this package
	"github.com/nex-gen-tech/nex/router"  // Assuming the router is in this package
)

const contentTypeHeader = "Content-Type"

// ContentTypeChecking ensures that the client sends requests with the expected content type.
func ContentTypeChecking(expectedContentType string) router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			// Check the Content-Type header of the incoming request
			contentType := c.Request.Header.Get(contentTypeHeader)

			// If the content type does not match the expected type, return an error
			if contentType != expectedContentType {
				http.Error(c.Response, "Invalid Content-Type", http.StatusUnsupportedMediaType)
				return
			}

			// Continue processing the request
			next(c)
		}
	}
}
