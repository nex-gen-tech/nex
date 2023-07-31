package nexmiddleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/nex-gen-tech/nex/context"
	"github.com/nex-gen-tech/nex/router"
)

const (
	// SecretKey used to sign JWTs. This should be kept private and secure.
	SecretKey = "your_secret_key_here"
)

// Authentication middleware using JWT
func Authentication() router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			// Extract token from Authorization header
			authHeader := c.Request.Header.Get("Authorization")
			if authHeader == "" {
				unauthorized(c, "Authorization header not provided")
				return
			}

			// The header should be in the format: Bearer <token>
			splitToken := strings.Split(authHeader, "Bearer ")
			if len(splitToken) != 2 {
				unauthorized(c, "Invalid Authorization header format")
				return
			}
			tokenString := splitToken[1]

			// Parse and validate the token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(SecretKey), nil
			})
			if err != nil {
				unauthorized(c, "Invalid token")
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// Add claims to context or do other authentication logic
				c.Params["userID"] = string(claims["userID"].(string))
				next(c)
			} else {
				unauthorized(c, "Invalid token")
				return
			}
		}
	}
}

// Helper function to respond with "Unauthorized"
func unauthorized(c *context.Context, message string) {
	http.Error(c.Response, message, http.StatusUnauthorized)
}
