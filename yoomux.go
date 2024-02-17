// Package yoomux provides a simple and flexible HTTP multiplexer with middleware support.
package yoomux

import (
	"net/http"
)

// Yoomux represents an HTTP multiplexer with middleware support.
type Yoomux struct {
	mux        *http.ServeMux
	middleware []func(http.Handler) http.Handler
}

// Use adds a middleware function to the middleware chain of the Yoomux instance.
// It returns a new Yoomux instance with the updated middleware chain.
func (yoomux *Yoomux) Use(middleware func(http.Handler) http.Handler) *Yoomux {
	return &Yoomux{
		mux:        yoomux.mux,
		middleware: append(yoomux.middleware, middleware),
	}
}

// UseAll adds a middleware function to the end of the middleware chain of the Yoomux instance.
// It modifies the existing Yoomux instance and returns it.
func (yoomux *Yoomux) UseAll(middleware func(http.Handler) http.Handler) *Yoomux {
	yoomux.middleware = append(yoomux.middleware, middleware)
	return yoomux
}

// Get registers a GET request handler function for the specified path.
// It applies the middleware chain to the handler function.
// It returns the Yoomux instance for method chaining.
func (yoomux *Yoomux) Get(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodGet+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

// Head registers a HEAD request handler function for the specified path.
// It applies the middleware chain to the handler function.
// It returns the Yoomux instance for method chaining.
func (yoomux *Yoomux) Head(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodHead+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

// Post registers a POST request handler function for the specified path.
// It applies the middleware chain to the handler function.
// It returns the Yoomux instance for method chaining.
func (yoomux *Yoomux) Post(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodPost+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

// Put registers a PUT request handler function for the specified path.
// It applies the middleware chain to the handler function.
// It returns the Yoomux instance for method chaining.
func (yoomux *Yoomux) Put(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodPut+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

// Patch registers a PATCH request handler function for the specified path.
// It applies the middleware chain to the handler function.
// It returns the Yoomux instance for method chaining.
func (yoomux *Yoomux) Patch(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodPatch+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

// Delete registers a DELETE request handler function for the specified path.
// It applies the middleware chain to the handler function.
// It returns the Yoomux instance for method chaining.
func (yoomux *Yoomux) Delete(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodDelete+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

// Connect registers a CONNECT request handler function for the specified path.
// It applies the middleware chain to the handler function.
// It returns the Yoomux instance for method chaining.
func (yoomux *Yoomux) Connect(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodConnect+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

// Options registers an OPTIONS request handler function for the specified path.
// It applies the middleware chain to the handler function.
// It returns the Yoomux instance for method chaining.
func (yoomux *Yoomux) Options(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodOptions+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

// Trace registers a TRACE request handler function for the specified path.
// It applies the middleware chain to the handler function.
// It returns the Yoomux instance for method chaining.
func (yoomux *Yoomux) Trace(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodTrace+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

// applyMiddleware applies the middleware chain to the handler function.
func (yoomux *Yoomux) applyMiddleware(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var h http.Handler = http.HandlerFunc(handler)
		for i := len(yoomux.middleware) - 1; i >= 0; i-- {
			h = yoomux.middleware[i](h)
		}
		h.ServeHTTP(w, r)
	}
}

// ServeHTTP implements the http.Handler interface.
// It delegates the request handling to the underlying http.ServeMux.
func (yoomux *Yoomux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	yoomux.mux.ServeHTTP(w, r)
}

// NewYoomux creates a new Yoomux instance with an empty middleware chain and a new http.ServeMux.
func NewYoomux() *Yoomux {
	return &Yoomux{
		mux:        http.NewServeMux(),
		middleware: make([]func(http.Handler) http.Handler, 0),
	}
}

// Subrouter creates a new Yoomux instance that acts as a subrouter for the specified path.
// It returns the subrouter Yoomux instance.
func (yoomux *Yoomux) Subrouter(path string) *Yoomux {
	submux := http.NewServeMux()
	yoomux.mux.Handle(path+"/", http.StripPrefix(path, submux))
	return &Yoomux{
		mux:        submux,
		middleware: yoomux.middleware,
	}
}
