package interceptor

import (
	"time"

	"github.com/nex-gen-tech/nex/context" // Assuming the context is now in this package
	"github.com/nex-gen-tech/nex/router"  // Assuming the router is in this package
)

const sessionCookieName = "SESSION_ID"

// Session represents a user's session with data.
type Session struct {
	ID        string
	UserID    string
	ExpiresAt time.Time
	Data      map[string]interface{}
}

// SessionStore represents a store for sessions.
type SessionStore interface {
	Get(sessionID string) (*Session, error)
	Save(session *Session) error
	Delete(sessionID string) error
}

// SessionManagement manages user sessions.
func SessionManagement(store SessionStore) router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			cookie, err := c.Request.Cookie(sessionCookieName)
			if err != nil || cookie.Value == "" {
				// No session cookie found, create a new session
				// ... (code to create a new session)
			} else {
				// Retrieve the session from the store
				session, err := store.Get(cookie.Value)
				if err != nil || session == nil || time.Now().After(session.ExpiresAt) {
					// Session not found or expired
					// ... (code to handle expired/missing session)
				} else {
					// Session found, add it to the context for further processing
					c.Set("session", session)
				}
			}

			// Continue processing the request
			next(c)
		}
	}
}
