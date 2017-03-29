// Package Hasty provides very simple and fast Multiplexer for Go
// Package is OpenSource at https://github.com/harshvladha/hasty
// Developed by Harsh Vardhan Ladha
// Email: harsh.ladha@gmail.com

package hasty

import "net/http"

// Route contains Path and methods on it is enabled,
// with Handler to serve request
// Path: is full path with prefix
// methods: have value OR'ed for all the methods enabled
// Handler: has a serving handler for the request
type Route struct {
	Path    string
	methods int
	methodHandlers map[string]*http.Handler
}

// NewRoute creates new instance of Route
// and returns its pointer
func NewRoute(url string) *Route {
	return &Route{Path: url, methodHandlers:make(map[string]*http.Handler)}
}

// setMethod enables the passed method for the given route
func (r *Route) setMethod(m string, handler *http.Handler) {
	r.methods |= methods[m]
	r.methodHandlers[m] = handler
}

// methodAllowed checks if Method is allowed on the route
func (r *Route) methodAllowed(m string) bool {
	return methods[m]&r.methods != 0
}

// handler takes method as argument and returns
// pointer of its applicable handler
func (r *Route) handler(m string) *http.Handler {
	return r.methodHandlers[m]
}