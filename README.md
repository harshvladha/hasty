hasty [![GoDoc](https://godoc.org/github.com/harshvladha/hasty?status.png)](http://godoc.org/github.com/harshvladha/hasty) [![Go Report Card](https://goreportcard.com/badge/harshvladha/hasty)](https://goreportcard.com/report/harshvladha/hasty) [![Sourcegraph](https://sourcegraph.com/github.com/harshvladha/hasty/-/badge.svg)](https://sourcegraph.com/github.com/harshvladha/hasty?badge)
=====
Hasty is Simple and Fast Multiplexer for Go language! It supports:
- URL Parameters
- Router Prefix
- Http method declaration
- Support for `http.Handler` and `http.HandlerFunc`
- Custom NotFound handler

### Example
``` go
package main

import (
	"net/http"
	"encoding/json"
	"github.com/harshvladha/hasty"
)

func main() {
	mux := hasty.New()
	
	// mux.Get, Post, Put, etc... takes http.Handler 
	mux.Get("/test", http.HandlerFunc(getHandler))
	
	// :var1 is a path variable
	mux.Get("/test/:var1", http.HandlerFunc(getHandler))
	mux.ListenAndServe("8080")
}

func getHandler(rw http.ResponseWriter, req *http.Request) {
	var queryId string
	// Get the value of var1 path variable
	if query := hasty.GetValue(req, "var1"); query != "" {
		queryId = query
	}
	var myJson = struct {
		First  string
		Second string
		Third  string
	}{"Hello", "World!", queryId}

	json.NewEncoder(rw).Encode(&myJson)
}

```

### TODO
- Test Cases
- REGEX support in Path Variable
- License
- Optimisation