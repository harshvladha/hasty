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

func (mux *Mux) Register(method string, path string, handler http.Handler) *Route {
	return mux.register(method, path, handler)
}

func (mux *Mux) GetFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodGet, path, handler)
}

func (mux *Mux) HeadFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodHead, path, handler)
}

func (mux *Mux) PostFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodPost, path, handler)
}

func (mux *Mux) PutFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodPut, path, handler)
}

func (mux *Mux) PatchFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodPatch, path, handler)
}

func (mux *Mux) DeleteFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodDelete, path, handler)
}

func (mux *Mux) ConnectFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodConnect, path, handler)
}

func (mux *Mux) OptionsFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodOptions, path, handler)
}

func (mux *Mux) TraceFunc(path string, handler http.HandlerFunc) *Route {
	return mux.register(http.MethodTrace, path, handler)
}

func (mux *Mux) Get(path string, handler http.Handler) *Route {
	return mux.register(http.MethodGet, path, handler)
}

func (mux *Mux) Head(path string, handler http.Handler) *Route {
	return mux.register(http.MethodHead, path, handler)
}

func (mux *Mux) Post(path string, handler http.Handler) *Route {
	return mux.register(http.MethodPost, path, handler)
}

func (mux *Mux) Put(path string, handler http.Handler) *Route {
	return mux.register(http.MethodPut, path, handler)
}

func (mux *Mux) Patch(path string, handler http.Handler) *Route {
	return mux.register(http.MethodPatch, path, handler)
}

func (mux *Mux) Delete(path string, handler http.Handler) *Route {
	return mux.register(http.MethodDelete, path, handler)
}

func (mux *Mux) Connect(path string, handler http.Handler) *Route {
	return mux.register(http.MethodConnect, path, handler)
}

func (mux *Mux) Options(path string, handler http.Handler) *Route {
	return mux.register(http.MethodOptions, path, handler)
}

func (mux *Mux) Trace(path string, handler http.Handler) *Route {
	return mux.register(http.MethodTrace, path, handler)
}

func (mux *Mux) HandleFunc(path string, handler http.HandlerFunc) {
	mux.Handle(path, handler)
}

func (mux *Mux) Handle(path string, handler http.Handler) {
	for m := range methods {
		mux.register(m, path, handler)
	}
}

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
