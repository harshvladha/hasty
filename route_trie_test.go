package hasty

import (
	"testing"
)

func TestSplitInTwo(t *testing.T) {
	a, b := SplitInTwo("my/dummy/path", "/")
	if a != "my" {
		t.Error("First Part doesn't matches")
	}

	if b != "dummy/path" {
		t.Error("Second Part doesn't matches")
	}

	a, b = SplitInTwo("mypath", "/")

	if b != "" {
		t.Error("Second part should be empty")
	}
}

func TestRootRouteTrie(t *testing.T) {
	rt := rootRouteTrie()
	if rt.token != "" {
		t.Error("Root Route is not an empty string")
	}
	if rt.pattern == true {
		t.Error("Root Route is pattern based")
	}
}

func TestNewRouteTrieWithoutPattern(t *testing.T) {
	rt := newRouteTrie("article")
	if rt.token != "article" {
		t.Error("Root trie path and token doesn't match")
	}
	if rt.pattern == true {
		t.Error("Pattern based route trie from non-pattern based path")
	}
}

func TestNewRouteTrieWithPattern(t *testing.T) {
	rt := newRouteTrie(":articleName")
	if rt.token != "articleName" {
		t.Error("Root trie path and token doesn't match")
	}
	if rt.pattern != true {
		t.Error("Not a pattern based route trie from pattern based path")
	}
}
