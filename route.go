/*
Package Hasty provides very simple and fast Multiplexer for Go
Package is OpenSource at https://github.com/harshvladha/hasty
Developed by Harsh Vardhan Ladha
Email: harsh.ladha@gmail.com
*/

package hasty

import "net/http"

type Route struct {
	Path    string
	methods int
	Handler http.Handler
}

func NewRoute(url string, handler http.Handler) *Route {
	return &Route{Path: url, Handler: handler}
}
