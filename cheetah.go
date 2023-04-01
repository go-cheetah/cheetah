package cheetah

import (
	"fmt"
	"net/http"
)

// Cheetah is a lightweight HTTP router.
type Cheetah struct {
	routes map[string]map[string]http.HandlerFunc
}

// New creates a new instance of Cheetah.
func New() *Cheetah {
	return &Cheetah{routes: make(map[string]map[string]http.HandlerFunc, 0)}
}

// ListenAndServe starts a new HTTP server listening on the given address.
func (c *Cheetah) ListenAndServe(addr string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: c,
	}
	return server.ListenAndServe()
}

// Get registers a new route with the GET HTTP method.
func (c *Cheetah) Get(path string, handlerFunc http.HandlerFunc) {
	c.Handle(http.MethodGet, path, handlerFunc)
}

// Post registers a new route with the POST HTTP method.
func (c *Cheetah) Post(path string, handlerFunc http.HandlerFunc) {
	c.Handle(http.MethodPost, path, handlerFunc)
}

// Put registers a new route with the PUT HTTP method.
func (c *Cheetah) Put(path string, handlerFunc http.HandlerFunc) {
	c.Handle(http.MethodPut, path, handlerFunc)
}

// Patch registers a new route with the PATCH HTTP method.
func (c *Cheetah) Patch(path string, handlerFunc http.HandlerFunc) {
	c.Handle(http.MethodPatch, path, handlerFunc)
}

// Delete registers a new route with the DELETE HTTP method.
func (c *Cheetah) Delete(path string, handlerFunc http.HandlerFunc) {
	c.Handle(http.MethodDelete, path, handlerFunc)
}

// Handle registers a new route with the given HTTP method and path.
func (c *Cheetah) Handle(httpMethod, path string, handlerFunc http.HandlerFunc) {
	if _, ok := c.routes[path]; !ok {
		c.routes[path] = make(map[string]http.HandlerFunc)
	}
	if _, ok := c.routes[path][httpMethod]; ok {
		panic(fmt.Sprintf("cheetah: handler already registered for %s %s", httpMethod, path))
	}
	c.routes[path][httpMethod] = handlerFunc
}

// ServeHTTP dispatches the HTTP request to the registered handler for the given route and HTTP method.
func (c *Cheetah) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, ok := c.routes[req.URL.Path][req.Method]; ok {
		handler(w, req)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
