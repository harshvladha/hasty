/*
Package Hasty provides very simple and fast Multiplexer for Go
Package is OpenSource at https://github.com/harshvladha/hasty
Developed by Harsh Vardhan Ladha
Email: harsh.ladha@gmail.com
*/

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
	Routes        map[string]*Route
	prefix        string
	notFound      http.Handler
	Serve         func(rw http.ResponseWriter, req *http.Request)
	caseSensitive bool
}

var (
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

func New() *Mux {
	mux := &Mux{Routes: make(map[string]*Route), caseSensitive: false}
	mux.Serve = mux.DefaultServe
	return mux
}

func (mux *Mux) CaseSensitive(cs bool) *Mux {
	mux.caseSensitive = cs
	return mux
}

func (mux *Mux) Prefix(prefix string) *Mux {
	mux.prefix = strings.TrimSuffix(prefix, "/")
	return mux
}

func (mux *Mux) DefaultServe(rw http.ResponseWriter, req *http.Request) {
	if ok, err := mux.validate(rw, req); !ok {
		rw.WriteHeader(err.HttpStatus)
	}
}

func (mux *Mux) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if !mux.caseSensitive {
		req.URL.Path = strings.ToLower(req.URL.Path)
	}
	mux.Serve(rw, req)
}
