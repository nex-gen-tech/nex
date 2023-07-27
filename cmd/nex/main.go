package main

import (
	"net/http"

	"github.com/nex-gen-tech/nex"
	"github.com/nex-gen-tech/nex/context"
	"github.com/nex-gen-tech/nexlog"
)

type User struct {
	Name  string `json:"name" nex:"required"`
	Email string `json:"email" nex:"required,email"`
	Age   int    `json:"age" nex:"required,gt=18"`
}

func main() {
	r := nex.New()

	// Define a handler
	handler := func(c *context.Context) {
		id := c.PathParam.Get("email")

		data := map[string]any{
			"message": "Hello " + "name : " + id,
		}

		c.Res.JSON(http.StatusOK, data)
	}

	// Register a route
	// r.GET("/hello/world/:id", handler)
	// A get path with params of regex which can match only email
	r.GET("/hello/world/:email([a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z0-9]+)", handler)

	// Create a new server
	logger := nexlog.New("New")

	logger.InfoF("Starting server on port %s", ":8080")

	r.Run(":8080")
}
