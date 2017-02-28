/*
Package Hasty provides very simple and fast Multiplexer for Go
Package is OpenSource at https://github.com/harshvladha/hasty
Developed by Harsh Vardhan Ladha
Email: harsh.ladha@gmail.com
*/

package hasty

import (
	"net/http"
)

// Register takes method, path and its handler to register the route
func (mux *Mux) Register(method string, path string, handler http.Handler) *Route {
	return mux.register(method, path, handler)
}

// GetFunc registers route with method GET
func (mux *Mux) GetFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodGet, path, handler)
}

// HeadFunc registers route with method HEAD
func (mux *Mux) HeadFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodHead, path, handler)
}

// PostFunc registers route with method POST
func (mux *Mux) PostFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodPost, path, handler)
}

// PutFunc registers route with method Put
func (mux *Mux) PutFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodPut, path, handler)
}

// PatchFunc registers route with method PATCH
func (mux *Mux) PatchFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodPatch, path, handler)
}

// DeleteFunc registers route with method DELETE
func (mux *Mux) DeleteFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodDelete, path, handler)
}

// ConnectFunc registers route with method CONNECT
func (mux *Mux) ConnectFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodConnect, path, handler)
}

// OptionsFunc registers route with method OPTIONS
func (mux *Mux) OptionsFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodOptions, path, handler)
}

// TraceFunc registers route with method TRACE
func (mux *Mux) TraceFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodTrace, path, handler)
}

// Get registers route with method GET
func (mux *Mux) Get(path string, handler http.Handler) *Route {
	return mux.register(http.MethodGet, path, handler)
}

// Head registers route with method HEAD
func (mux *Mux) Head(path string, handler http.Handler) *Route {
	return mux.register(http.MethodHead, path, handler)
}

// Post registers route with method POST
func (mux *Mux) Post(path string, handler http.Handler) *Route {
	return mux.register(http.MethodPost, path, handler)
}

// Put registers route with method PUT
func (mux *Mux) Put(path string, handler http.Handler) *Route {
	return mux.register(http.MethodPut, path, handler)
}

// Patch registers route with method PATCH
func (mux *Mux) Patch(path string, handler http.Handler) *Route {
	return mux.register(http.MethodPatch, path, handler)
}

// Delete registers route with method DELETE
func (mux *Mux) Delete(path string, handler http.Handler) *Route {
	return mux.register(http.MethodDelete, path, handler)
}

// Connect registers route with method CONNECT
func (mux *Mux) Connect(path string, handler http.Handler) *Route {
	return mux.register(http.MethodConnect, path, handler)
}

// Options registers route with method OPTIONS
func (mux *Mux) Options(path string, handler http.Handler) *Route {
	return mux.register(http.MethodOptions, path, handler)
}

// Trace registers route with method TRACE
func (mux *Mux) Trace(path string, handler http.Handler) *Route {
	return mux.register(http.MethodTrace, path, handler)
}

// HandleFunc registers route with all the methods
func (mux *Mux) HandleFunc(path string, handler http.HandlerFunc) {
	mux.Handle(path, handler)
}

// Handle registers  route with all the methods
func (mux *Mux) Handle(path string, handler http.Handler) {
	for m := range methods {
		mux.register(m, path, handler)
	}
}

// register cleans URL and registers the method for the given route
// if route doesn't already exist, it creates new Route instance
func (mux *Mux) register(method string, path string, handler http.Handler) *Route {
	fullUrl := mux.prefix + path
	cleanURL(&fullUrl)
	route := mux.Routes[fullUrl]
	if route == nil {
		route = NewRoute(fullUrl, handler)
		mux.Routes[fullUrl] = route
	}
	route.methods |= methods[method]
	return route
}
