# Cheetah
[![GoDoc](https://img.shields.io/badge/go-documentation-blue?style=flat-square)](https://pkg.go.dev/github.com/go-cheetah/cheetah)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-cheetah/cheetah?style=flat-square)](https://goreportcard.com/report/github.com/go-cheetah/cheetah)
[![Codecov](https://img.shields.io/codecov/c/github/go-cheetah/cheetah?style=flat-square)](https://codecov.io/gh/go-cheetah/cheetah)
[![License](https://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/go-cheetah/cheetah/main/LICENSE)

Cheetah is a lightweight HTTP router for Go, designed to be simple and easy to use.

## Installation
You can install Cheetah using Go modules:

``` zsh
go get -u github.com/go-cheetah/cheetah
```

## Usage
Here's a simple example that shows how to use Cheetah to create a simple web server:
``` go
package main

import (
	"fmt"
	"net/http"

	"github.com/go-cheetah/cheetah"
)

func main() {
	c := cheetah.New()

	c.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	if err := c.ListenAndServe(":3000"); err != nil {
        panic(err)
    }
}
```
In this example, we create a new Cheetah server, register a GET route, and start the server listening on port 3000.

Cheetah provides a similar API to the standard library http package, with methods for handling GET, POST, PUT, PATCH, and DELETE requests. You can also use the Handle method to register a handler for any HTTP method.

``` go
router.Post("/users", func(w http.ResponseWriter, r *http.Request) {
	// Handle POST request to /users
})

router.Handle("OPTIONS", "/users", func(w http.ResponseWriter, r *http.Request) {
	// Handle OPTIONS request to /users
})
```

## Documentation
For more information about Cheetah and its API, see the GoDoc documentation.

## Contributing
Contributions to Cheetah are always welcome! If you find a bug, or have a feature request, please open an issue. If you'd like to contribute code, please open a pull request.

## License
Cheetah is licensed under the MIT License.