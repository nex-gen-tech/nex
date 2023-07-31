package middleware

import (
	"fmt"

	"github.com/nex-gen-tech/nex/context" // Assuming the context is now in this package
	"github.com/nex-gen-tech/nex/router"  // Assuming the router is in this package
)

// Logging logs the method, path, request ID, IP, user agent, and time for each request.
// func Logging() router.MiddlewareFunc {
// 	logOutput := log.New(os.Stdout, "NEX-LOG", 0)

// 	return func(next router.HandlerFunc) router.HandlerFunc {
// 		return func(c *context.Context) {
// 			start := time.Now()

// 			fmt.Println("Logging middleware initialized")

// 			// Call the next handler
// 			next(c)

// 			// Extract details from the request
// 			timestamp := start.Format(time.RFC3339) // use the start time as the timestamp
// 			method := c.Request.Method
// 			reqID := c.Params["requestID"] // assuming the requestID is saved as a route parameter
// 			path := c.Request.URL.Path
// 			ip := c.Request.RemoteAddr

// 			// Calculate duration
// 			duration := time.Since(start)

// 			// Log the details
// 			logOutput.Printf("%s %s %s %s %s %s \n",
// 				color.New(color.BgHiGreen).Sprintf(" %s ", method),
// 				color.New(color.FgHiGreen).Sprintf(" %s ", reqID),
// 				timestamp,
// 				path,
// 				duration,
// 				ip,
// 			)
// 		}
// 	}
// }

func Logging() router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			fmt.Println("Route middleware initialized")
			next(c)
			fmt.Println("Route middleware exited")
		}
	}
}

func LoggingGroup() router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			fmt.Println("Group middleware initialized")
			next(c)
			fmt.Println("Group middleware exited")
		}
	}
}

func LoggingRouter() router.MiddlewareFunc {
	return func(next router.HandlerFunc) router.HandlerFunc {
		return func(c *context.Context) {
			fmt.Println("Router middleware initialized")
			next(c)
			fmt.Println("Router middleware exited")
		}
	}
}
