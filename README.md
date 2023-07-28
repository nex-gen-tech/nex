# Nex Web Framework

Nex is a lightweight, high-performance web framework for Go. Designed with simplicity and scalability in mind, Nex provides a robust set of features for building web applications and APIs.

![Nex Logo](./assets/nex-logo.png)

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Middleware](#middleware)
- [Routing](#routing)
- [Grouping Routes](#grouping-routes)
- [Error Handling](#error-handling)
- [Best Practices](#best-practices)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Fast Routing**: Efficient routing using a trie-based structure.
- **Middleware Support**: Easily extend functionality with middleware.
- **Grouping Routes**: Organize your routes with groups.
- **Context**: A comprehensive context for handlers, making it easier to manage request and response data.
- **Error Handling**: Built-in error handling mechanism.
- **Session Management**: Secure and efficient session management capabilities.
- **Compression**: Automatic response compression.
- **Rate Limiting**: Protect your application from abuse with rate limiting.
- **CORS**: Built-in support for Cross-Origin Resource Sharing.
- **Cache Control**: Efficiently manage cache headers and server-side caching.
- **Content-Type Checking**: Ensure that your endpoints receive the expected content types.

## Installation

To install Nex, use `go get`:

```bash
go get github.com/nex-gen-tech/nex
```

## Quick Start

Here's a simple example to get you started:

```go
package main

import (
	"github.com/nex-gen-tech/nex"
	"github.com/nex-gen-tech/nex/context"
)

func main() {
	r := nex.NewRouter()

	r.GET("/", func(c *context.Context) {
		c.String(200, "Welcome to Nex!")
	})

	r.Run(":8080")
}
```

## Middleware

Nex supports middleware at both the global and group levels. Here's how you can use the built-in logging and recovery middleware:

```go
r := nex.NewRouter()

r.Use(middleware.Logging(), middleware.Recovery())

// ... your routes here ...
```

## Routing

Define routes easily with Nex:

```go
r.GET("/users", listUsers)
r.POST("/users", createUser)
```

## Grouping Routes

Organize your routes with groups:

```go
api := r.NewGroup("/api")

api.GET("/users", listUsers)
api.POST("/users", createUser)
```

## Error Handling

Nex provides a built-in mechanism for error handling:

```go
r.SetErrorHandler(func(c *context.Context, err error) {
	// Handle the error here
	c.String(500, "Internal Server Error")
})
```

## Best Practices

- Always check for errors and handle them gracefully.
- Use middleware to extend functionality and keep your code DRY.
- Organize your routes with groups, especially for larger applications.
- Monitor your application's performance and adjust rate limits, compression settings, etc., as needed.

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](./CONTRIBUTING.md) for details on how to contribute.

## License

Nex is licensed under the MIT License. See [LICENSE](./LICENSE) for more information.
