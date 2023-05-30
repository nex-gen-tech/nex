package middleware

import (
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/nex-gen-tech/nex"
)

// Logging logs the method, path, status, request ID, IP, user agent, and time for each request.
func Logging() nex.MiddlewareFunc {
	log := log.New(os.Stdout, "[NEX] ", 0)

	return func(next nex.HandlerFunc) nex.HandlerFunc {
		return func(c *nex.Context) {
			start := time.Now()

			next(c)

			// Extract details from the request
			timestamp := start.Format(time.DateTime) // use the start time as the timestamp
			method := c.Req.Method
			status := c.Get("status")            // assuming the status is saved in the context
			reqID := c.Get("requestID").(string) // assuming the status is saved in the context
			path := c.Req.URL.Path
			ip := c.Req.RemoteAddr
			// errorMsg := c.Get("error").(string) // assuming the error is saved in the context
			// Calculate duration
			duration := time.Since(start)

			// Log the details
			log.Printf("%s %s %s %s %s %s %s \n",
				color.New(color.BgHiGreen).Sprintf(" %s ", method),
				color.New(color.BgHiGreen).Sprintf(" %s ", status),
				color.New(color.FgHiGreen).Sprintf(" %s ", reqID),
				timestamp,
				path,
				duration,
				ip,
			)
		}
	}
}

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
