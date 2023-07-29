package main

import (
	"fmt"
	"net/http"

	"github.com/nex-gen-tech/nex"
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
	handler := func(c *nex.Context) {
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

	// Api Group
	api := r.Group("/api")

	// get user
	api.GET("/user/:id", func(c *nex.Context) {
		id := c.PathParam.Get("id")

		// get is_active query param
		isActive, err := c.QueryParam.GetAsBool("is_active")
		if err != nil {
			c.Res.JsonBadRequest400("is_active query param is not a boolean")
			return
		}

		data := map[string]any{
			"message":   "Hello " + "name : " + id,
			"is_active": isActive,
		}

		c.Res.JsonOk200(data, "user fetched successfully")
	})

	// create user
	api.POST("/user", func(c *nex.Context) {
		var user User
		if err := c.Body.ParseJSON(&user); err != nil {
			c.Res.JsonBadRequest400("invalid json body")
			return
		}

		// validate user
		if err := c.Validate(&user); err != nil {
			c.Res.JsonBadRequest400(fmt.Sprintf("invalid user: %v", err[0]))
			return
		}

		c.Res.JsonOk200(user, "user created successfully")
	})

	// Create a new server
	logger := nexlog.New("NEX-TEST")

	logger.InfoF("Starting server on port %s", ":8080")

	r.Run(":8080")
}
