package middleware

import (
	"log"

	"github.com/nex-gen-tech/nex"
)

// Recovery recovers from panics and logs a stack trace.
func Recovery() nex.MiddlewareFunc {
	return func(next nex.HandlerFunc) nex.HandlerFunc {
		return func(c *nex.Context) {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("Panic: %v", err)
				}
			}()
			next(c)
		}
	}
}
