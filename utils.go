/*
Package Hasty provides very simple and fast Multiplexer for Go
Package is OpenSource at https://github.com/harshvladha/hasty
Developed by Harsh Vardhan Ladha
Email: harsh.ladha@gmail.com
*/

package hasty

import "net/http"

type ErrorStatusCode struct {
	HttpStatus int
}

func (mux *Mux) ListenAndServe(port string) error {
	return http.ListenAndServe(":"+port, mux)
}

func (mux *Mux) validate(rw http.ResponseWriter, req *http.Request) (bool, ErrorStatusCode) {
	pathLen := len(req.URL.Path)
	if pathLen > 1 && req.URL.Path[pathLen-1:] == "/" {
		cleanURL(&req.URL.Path)
		rw.Header().Set("Location", req.URL.String())
		rw.WriteHeader(http.StatusFound)
		return true, nil
	}
	return mux.parse(rw, req)
}

func cleanURL(url *string) {
	urlLen := len(*url)
	if urlLen > 1 {
		if (*url)[urlLen-1:] == "/" {
			*url = (*url)[:urlLen-1]
			cleanURL(url)
		}
	}
}

func (mux *Mux) parse(rw http.ResponseWriter, req *http.Request) (bool, ErrorStatusCode) {
	if mux.Routes[req.URL.Path] == nil {
		return false, ErrorStatusCode{HttpStatus: http.StatusNotFound}
	}
	requestMethod := req.Method
	// for HEAD method, default to GET
	if requestMethod == http.MethodHead {
		requestMethod = http.MethodGet
	}
	// check if Method is allowed for the given route
	if (methods[requestMethod])&(mux.Routes[req.URL.Path].methods) == 0 {
		return false, ErrorStatusCode{HttpStatus: http.StatusMethodNotAllowed}
	}
	return true, nil
}
