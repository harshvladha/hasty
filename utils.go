/*
Package Hasty provides very simple and fast Multiplexer for Go
Package is OpenSource at https://github.com/harshvladha/hasty
Developed by Harsh Vardhan Ladha
Email: harsh.ladha@gmail.com
*/

package hasty

import "net/http"

// ErrorStatusCode is used to signal for
// HTTP Error Status Codes
type ErrorStatusCode struct {
	HttpStatus int
}

// ListenAndServe appends ":" in port before
// calling http.ListenAndServe
func (mux *Mux) ListenAndServe(port string) error {
	return http.ListenAndServe(":"+port, mux)
}

// validate is used to check for trailing slash
// and parse the request
func (mux *Mux) validate(rw http.ResponseWriter, req *http.Request) (bool, *ErrorStatusCode) {
	pathLen := len(req.URL.Path)
	if pathLen > 1 && req.URL.Path[pathLen-1:] == "/" {
		cleanURL(&req.URL.Path)
		rw.Header().Set("Location", req.URL.String())
		rw.WriteHeader(http.StatusFound)
		return true, nil
	}
	return mux.parse(rw, req)
}

// cleanURL cleans trailing slashes recursively
func cleanURL(url *string) {
	urlLen := len(*url)
	if urlLen > 1 {
		if (*url)[urlLen-1:] == "/" {
			*url = (*url)[:urlLen-1]
			cleanURL(url)
		}
	}
}

// parse checks for the route and its method
// returns http.StatusNotFound if route is not registered
// returns http.StatusMethodNotAllowed if method for that route is not registered
// otherwise calls the ServeHTTP of the http.Handler registered for the route
func (mux *Mux) parse(rw http.ResponseWriter, req *http.Request) (bool, *ErrorStatusCode) {
	if mux.Routes[req.URL.Path] == nil {
		return false, &ErrorStatusCode{HttpStatus: http.StatusNotFound}
	}
	requestMethod := req.Method
	// for HEAD method, default to GET
	if requestMethod == http.MethodHead {
		requestMethod = http.MethodGet
	}
	if !mux.Routes[req.URL.Path].methodAllowed(requestMethod) {
		return false, &ErrorStatusCode{HttpStatus: http.StatusMethodNotAllowed}
	}
	mux.Routes[req.URL.Path].Handler.ServeHTTP(rw, req)
	return true, nil
}
