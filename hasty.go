// Package Hasty provides very simple and fast Multiplexer for Go
// Package is OpenSource at https://github.com/harshvladha/hasty
// Developed by Harsh Vardhan Ladha
// Email: harsh.ladha@gmail.com

package hasty

import (
	"net/http"
	"strings"
)

// Mux have routes, notFound (404 status code) handler and Serve which serves incoming requests
// Routes: all the registered routes
// notFound: 404 status code handler, default: http.NotFound
// CaseSensitive: for routes to be case sensitive or not, default false
type Mux struct {
	Routes        *RouteTrie
	prefix        string
	notFound      http.Handler
	Serve         func(rw http.ResponseWriter, req *http.Request)
	caseSensitive bool
}

var (
	// methods is a map of HTTP Methods with a value 2^(x)
	// to determine methods enabled for
	// the given endpoint
	methods = map[string]int{
		http.MethodGet:     1 << 1,
		http.MethodHead:    1 << 2,
		http.MethodPost:    1 << 3,
		http.MethodPut:     1 << 4,
		http.MethodPatch:   1 << 5,
		http.MethodDelete:  1 << 6,
		http.MethodConnect: 1 << 7,
		http.MethodOptions: 1 << 8,
		http.MethodTrace:   1 << 9,
	}
)

// New creates new instance of Mux and returns its pointer
func New() *Mux {
	mux := &Mux{Routes: rootRouteTrie(), caseSensitive: false}
	mux.Serve = mux.DefaultServe
	return mux
}

// CaseSensitive sets whether routes are case sensitive
// or not. Default being false
func (mux *Mux) CaseSensitive(cs bool) *Mux {
	mux.caseSensitive = cs
	return mux
}

// Prefix sets a default prefix to all the registered routes
func (mux *Mux) Prefix(prefix string) *Mux {
	cleanURL(&prefix)
	mux.prefix = prefix
	return mux
}

// DefaultServe is the default HTTP request handler
func (mux *Mux) DefaultServe(rw http.ResponseWriter, req *http.Request) {
	ok, err := mux.validate(rw, req)
	if !ok {
		rw.WriteHeader(err.HttpStatus)
		if err.HttpStatus == http.StatusNotFound {
			mux.notFound.ServeHTTP(rw, req)
		}
	}
}

// ServeHTTP calls the Serve method of Mux for processing
func (mux *Mux) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if !mux.caseSensitive {
		req.URL.Path = strings.ToLower(req.URL.Path)
	}
	mux.Serve(rw, req)
}
