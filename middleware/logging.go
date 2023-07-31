package nexmiddleware

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/nex-gen-tech/nex/context"
	"github.com/nex-gen-tech/nex/router"
)

// Assuming the router is in this package

// Logging logs the method, path, request ID, IP, user agent, and time for each request.
func Logging() router.MiddlewareFunc {
	logOutput := log.New(os.Stdout, "NEX-LOG", 0)

	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			start := time.Now()

			fmt.Println("Logging middleware initialized")

			// Call the next handler
			next(c)

			// Extract details from the request
			timestamp := start.Format(time.RFC3339) // use the start time as the timestamp
			method := c.Request.Method
			reqID := c.Params["requestID"] // assuming the requestID is saved as a route parameter
			path := c.Request.URL.Path
			ip := c.Request.RemoteAddr

			// Calculate duration
			duration := time.Since(start)

			// Log the details
			logOutput.Printf("%s %s %s %s %s %s \n",
				color.New(color.BgHiGreen).Sprintf(" %s ", method),
				color.New(color.FgHiGreen).Sprintf(" %s ", reqID),
				timestamp,
				path,
				duration,
				ip,
			)
		}
	}
}
