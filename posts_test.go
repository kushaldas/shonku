package main

import (
	"io/ioutil"
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
		dat, _ := ioutil.ReadFile("./posts/hello-world.md")
		lines := strings.Split(string(dat), "\n")
		if lines[0] != "title: hello world" {
			t.Error("Title of the post is not matching.", lines[0])
		}
		if !strings.HasPrefix(lines[1], "date:") {
			t.Error("Date of the post is not matching.", lines[1])
		}
		if !strings.HasPrefix(lines[2], "tags: Blog") {
			t.Error("Tags of the post is not matching.", lines[2])
		}
		if lines[3] != "====" {
			t.Error("End of the metadata is not matching.", lines[3])
		}
	} else {
		t.Error("Missing new post")
	}
}
