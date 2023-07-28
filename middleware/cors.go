package middleware

import (
	"net/http"
	"strings"

	"github.com/nex-gen-tech/nex/context"
	"github.com/nex-gen-tech/nex/router"
)

// CORSConfig defines the config for CORS middleware.
type CORSConfig struct {
	AllowOrigins []string
	AllowMethods []string
	AllowHeaders []string
}

// DefaultCORSConfig is a basic default configuration for CORS.
var DefaultCORSConfig = CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
	AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
}

// CORS middleware
func CORS(config CORSConfig) router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			origin := c.Request.Header.Get("Origin")

			// Check if the origin is allowed
			if isOriginAllowed(origin, config.AllowOrigins) {
				c.Response.Header().Set("Access-Control-Allow-Origin", origin)
				c.Response.Header().Set("Access-Control-Allow-Methods", strings.Join(config.AllowMethods, ","))
				c.Response.Header().Set("Access-Control-Allow-Headers", strings.Join(config.AllowHeaders, ","))
			}

			// If it's a preflight request, respond with 200 OK
			if c.Request.Method == "OPTIONS" {
				c.Response.WriteHeader(http.StatusOK)
				return
			}

			next(c)
		}
	}
}

// Helper function to check if an origin is allowed
func isOriginAllowed(origin string, allowedOrigins []string) bool {
	for _, o := range allowedOrigins {
		if o == "*" || o == origin {
			return true
		}
	}
	return false
}
