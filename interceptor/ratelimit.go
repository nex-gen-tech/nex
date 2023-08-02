package interceptor

import (
	"net/http"
	"time"

	"github.com/nex-gen-tech/nex/context"
	"github.com/nex-gen-tech/nex/router"
)

const (
	// MaxRequests is the number of allowed requests in the defined time window.
	MaxRequests = 100

	// TimeWindow is the duration in which the requests are counted.
	TimeWindow = time.Hour
)

// RateLimiter is a middleware for rate limiting requests.
func RateLimiter(store Store) router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			// Identify user, for simplicity we're using IP. Consider using a more robust identification method for production.
			identifier := c.Request.RemoteAddr

			// Fetch current request count for user
			count, err := store.GetRequestCount(identifier)
			if err != nil {
				http.Error(c.Response, "Server error", http.StatusInternalServerError)
				return
			}

			// Check if user has exceeded their limit
			if count >= MaxRequests {
				http.Error(c.Response, "Rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			// Increment the request count
			err = store.IncrementRequestCount(identifier, TimeWindow)
			if err != nil {
				http.Error(c.Response, "Server error", http.StatusInternalServerError)
				return
			}

			// Continue processing the request
			next(c)
		}
	}
}

// Store interface defines methods to interact with the backend store.
type Store interface {
	GetRequestCount(identifier string) (int, error)
	IncrementRequestCount(identifier string, duration time.Duration) error
}

// For the sake of this example, we've defined a Store interface. In a real-world scenario,
// you'd implement this interface with a fast storage solution like Redis.
