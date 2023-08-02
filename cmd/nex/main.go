package main

import (
	"log"

	"github.com/nex-gen-tech/nex"
	"github.com/nex-gen-tech/nex/interceptor"

	"github.com/nex-gen-tech/nexlog"
)

type User struct {
	Name  string `json:"name" nex:"required"`
	Email string `json:"email" nex:"required,email"`
	Age   int    `json:"age" nex:"required,gt=18"`
}

func main() {
	r := nex.New()

	r.Use(interceptor.TestOne())

	group1 := r.Group("/group1")
	group1.Use(interceptor.TestTwo())

	group2 := group1.Group("/group2")
	group2.Use(interceptor.TestThree())

	group3 := group2.Group("/group3")
	group3.Use(interceptor.TestFour())

	group4 := group3.Group("/group4")
	group4.Use(interceptor.TestFive())

	group4.GET("/test", func(c *nex.Context) {
		c.Res.JsonOk200(nil, "Hello from test route")
	})

	// Create a new server
	logger := nexlog.New("NEX-TEST")

	logger.InfoF("Starting server on port %s", ":8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
