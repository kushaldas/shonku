package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var input io.Reader = os.Stdin
var pretext string = `<!--
.. title: %s
.. slug: %s
.. date: %s
.. tags: Blog
.. link:
.. description:
.. type: text
-->

Write your post here.
`

/*
Creates a new post file.
*/
func new_post() {
	const longform = "2006/01/02 15:04:05 MST"
	var title string
	fmt.Print("Enter the title of the post: ")
	in := bufio.NewReader(input)
	title, _ = in.ReadString('\n')
	title = strings.TrimSpace(title)
	slug := get_slug(title)
	name := "./posts/" + slug + ".md"
	f, err := os.Create(name)
	defer f.Close()
	if err == nil {
		t := time.Now()
		text := fmt.Sprintf(pretext, title, slug, t.Format(longform))
		f.WriteString(text)
		fmt.Println("Your new post is ready at " + name)
	}

}

/*
TODO: Creating new page should be updated in future.
*/
func new_page() {
	const longform = "2006/01/02 15:04:05 MST"
	var title string
	fmt.Print("Enter the title of the page: ")
	in := bufio.NewReader(input)
	title, _ = in.ReadString('\n')
	title = strings.TrimSpace(title)
	slug := get_slug(title)
	name := "./pages/" + slug + ".md"
	f, err := os.Create(name)
	defer f.Close()
	if err == nil {
		t := time.Now()
		text := fmt.Sprintf(pretext, title, slug, t.Format(longform))
		f.WriteString(text)
		fmt.Println("Your new page is ready at " + name)

	}

}
