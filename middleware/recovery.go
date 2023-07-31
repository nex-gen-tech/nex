package nexmiddleware

import (
	"log"

	"github.com/nex-gen-tech/nex/context" // Assuming the context is now in this package
	"github.com/nex-gen-tech/nex/router"  // Assuming the router is in this package
)

// Recovery recovers from panics and logs a stack trace.
func Recovery() router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("Panic: %v", err)
				}
			}()
			next(c)
		}
	}
}
