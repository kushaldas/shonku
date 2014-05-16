package main

import (
	"os"
	"strings"
	"testing"
)

func texists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Test_CreatePost(t *testing.T) {
	input = strings.NewReader("hello world\n")
	new_post()
	if texists("./posts/hello-world.md") {
	} else {
		t.Error("Missing new post")
	}
}
