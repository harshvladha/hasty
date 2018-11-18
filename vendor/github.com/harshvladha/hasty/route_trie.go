// Package Hasty provides very simple and fast Multiplexer for Go
// Package is OpenSource at https://github.com/harshvladha/hasty
// Developed by Harsh Vardhan Ladha
// Email: harsh.ladha@gmail.com

package hasty

import (
	"fmt"
	"io"
	"net/http"
)

// RouteTrie is a Trie data structure which holds
// route mapping with path variables
// it also avoids creating new RouteTrie objects for
// same parent RouteTrie (repetition)
type RouteTrie struct {
	token    string
	pattern  bool
	children []*RouteTrie
	route    *Route
}

// add adds the path to RouteTrie after scanning
// for parents, during Route Registration
func (rt *RouteTrie) add(path string) *RouteTrie {
	return rt.findOrCreate(path)
}

// findOrCreate tries to find the parent first
// if there is no parent it creates it and adds
// it to #RouteTrie.children slice
func (rt *RouteTrie) findOrCreate(path string) *RouteTrie {
	cleanURL(&path)
	current, next := SplitInTwo(path, "/")
	child := rt.findChild(current)
	if child == nil {
		child = newRouteTrie(current)
		rt.children = append(rt.children, child)
	}

	if len(next) > 0 {
		return child.findOrCreate(next)
	}

	return child
}

// findChild finds and returns child with the
// given token value
func (rt *RouteTrie) findChild(token string) *RouteTrie {
	if token[:1] == ":" {
		token = token[1:]
	}
	for _, child := range rt.children {
		if child.token == token {
			return child
		}
	}

	return nil
}

// parse parses the incoming request, process
// path variables, and return status codes
// if given route doesn't exists
func (rt *RouteTrie) parse(mux *Mux, path string, pathVars map[string]string) (*Route, *ErrorStatusCode) {
	cleanURL(&path)
	current, next := SplitInTwo(path, "/")
	child, matched := rt.matchAndParse(mux, current, pathVars)
	if matched {
		switch {
		case len(next) > 0:
			return child.parse(mux, next, pathVars)
		case child.route != nil:
			return child.route, nil
		}
	}

	return nil, &ErrorStatusCode{HttpStatus: http.StatusNotFound}
}

// matchAndParse finds applicable child RouteTries nodes,
// checks if the given token matches with RouteTrie token
// and saves path variable for pattern based RouteTrie node
func (rt *RouteTrie) matchAndParse(mux *Mux, path string, pathVars map[string]string) (*RouteTrie, bool) {
	for _, child := range rt.children {
		switch {
		case !child.pattern && child.token == path:
			return child, true
		case child.pattern:
			pathVars[child.token] = path
			return child, true
		}
	}

	return nil, false
}

// rootRouteTrie returns top most level RouteTrie
func rootRouteTrie() *RouteTrie {
	return &RouteTrie{token: "", pattern: false}
}

// newRouteTrie return new RouteTrie with given token
// checks if given token is of pattern type or not
func newRouteTrie(token string) *RouteTrie {
	pattern := false

	if token[:1] == ":" {
		pattern = true
		token = token[1:]
	}

	return &RouteTrie{token: token, pattern: pattern}
}

// SplitInTwo splits the given path in two parts
// using separator, it scans based on first encounter
// of the separator
func SplitInTwo(str string, sep string) (string, string) {
	idx := 0
	for i, c := range str {
		if string(c) == sep {
			idx = i
			break
		}
	}

	if idx == 0 {
		return str, ""
	}

	return str[:idx], str[idx+1:]
}

// String prints the RouteTrie graph
func (rt *RouteTrie) String(w io.Writer) {
	fmt.Fprint(w, rt.token+"[")
	for _, child := range rt.children {
		child.String(w)
	}
	fmt.Fprint(w, "]")
}
