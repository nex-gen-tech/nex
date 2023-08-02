package interceptor

import (
	"net/http"

	"github.com/nex-gen-tech/nex/context"
	"github.com/nex-gen-tech/nex/router"
)

// Define roles and their permissions
var rolePermissions = map[string][]string{
	"Admin": {"read", "write", "delete"},
	"User":  {"read", "write"},
	"Guest": {"read"},
}

// Authorization middleware
func Authorization(requiredPermission string) router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			// Fetch the user's role. This could be part of the JWT or fetched from a database.
			// For this example, we'll assume it's part of the JWT claims.
			userRole := c.Params["role"]

			// Check if the user's role has the required permission
			if hasPermission(userRole, requiredPermission) {
				next(c)
			} else {
				forbidden(c, "You don't have the required permissions to access this resource")
			}
		}
	}
}

// Helper function to check if a role has a specific permission
func hasPermission(role string, permission string) bool {
	permissions, exists := rolePermissions[role]
	if !exists {
		return false
	}

	for _, perm := range permissions {
		if perm == permission {
			return true
		}
	}
	return false
}

// Helper function to respond with "Forbidden"
func forbidden(c *context.Context, message string) {
	http.Error(c.Response, message, http.StatusForbidden)
}
