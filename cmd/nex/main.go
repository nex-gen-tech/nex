package main

import (
	"net/http"

	"github.com/nex-gen-tech/nex"
	"github.com/nex-gen-tech/nex/middleware"
	"github.com/nex-gen-tech/nexlog"
)

func GetUserIdAuth(c *nex.Context) {
	data := map[string]any{
		"id":     c.ParamMustGetInt("id"),
		"authId": c.ParamMustGetInt("authId"),
	}

	c.ResponseJson(http.StatusOK, data)
}

func GetUserId(c *nex.Context) {
	id, err := c.Param.GetParam("id").AsInt()
	if err != nil {
		c.ResNexBadRequest400(err.Error(), nil)
		return
	}

	data := map[string]any{
		"id": id,
	}

	c.ResponseJson(http.StatusOK, data)
}

type User struct {
	Name  string `json:"name" nex:"required"`
	Email string `json:"email" nex:"required,email"`
	Age   int    `json:"age" nex:"required,gt=18"`
}

// CreateUser - creates a new user.
func CreateUser(c *nex.Context) {
	var user User
	if err := c.BodyBindJson(&user); err != nil {
		c.ResNexBadRequest400(err.Error(), nil)
		return
	}

	// Validate the user
	if err := c.Validate(&user); err != nil {
		c.ResNexBadRequest400("Validation failed.", err)
		return
	}

	c.ResNexCreated201(user, "User created successfully.")
}

func main() {
	r := nex.New()

	r.Use(middleware.Logging())
	r.Use(middleware.Recovery())

	// Define a handler
	handler := func(c *nex.Context) {
		// name, err := c.Query.GetQuery("name").WithRequired().AsString()
		// if err != nil {
		// 	c.ResNexBadRequest400(err.Error(), nil)
		// 	return
		// }

		data := map[string]any{
			"message": "Hello " + "name",
		}

		c.ResponseJsonOk200(data)
	}

	// Register a route
	r.GET("/hello/world", handler)
	// r.GET("/hello/world", handler)
	r.GET("/user/:id/:authId", GetUserIdAuth)
	r.GET("/user/:id", GetUserId)

	api := r.Group("api")
	api.GET("/rama", handler)

	v1 := api.Group("v1")

	// for i := 1; i <= 1000; i++ {
	// 	v1.GET(fmt.Sprintf("/user/%d", i), handler)
	// }

	v1.POST("/user", CreateUser)
	user := v1.Group("user")
	{
		user.GET("/get", handler)
	}

	api.PrintRoutes(&nex.PrintRouteConfig{WriteToFile: true})

	// Create a new server
	logger := nexlog.New("New")

	logger.InfoF("Starting server on port %s", ":8080")

	r.Run(":8080")
}
