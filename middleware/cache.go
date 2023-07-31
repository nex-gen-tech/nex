package nexmiddleware

import (
	"strconv"
	"time"

	"github.com/nex-gen-tech/nex/context" // Assuming the context is now in this package
	"github.com/nex-gen-tech/nex/router"  // Assuming the router is in this package
)

// CacheControl sets the Cache-Control header for the response.
func CacheControl(maxAge time.Duration, directives ...string) router.MiddlewareFunc {
	cacheControlValue := "max-age=" + strconv.Itoa(int(maxAge.Seconds()))
	for _, directive := range directives {
		cacheControlValue += ", " + directive
	}

	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			c.Response.Header().Set("Cache-Control", cacheControlValue)
			next(c)
		}
	}
}
