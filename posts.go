package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func new_post() {
	const longform = "2006/01/02 15:04:05"
	var title string
	fmt.Print("Enter the title of the post: ")
	in := bufio.NewReader(os.Stdin)
	title, _ = in.ReadString('\n')
	title = strings.TrimSpace(title)
	slug := get_slug(title)
	name := "./posts/" + slug + ".md"
	f, err := os.Create(name)
	defer f.Close()
	if err == nil {
		f.WriteString("title: " + title + "\n")
		t := time.Now()
		f.WriteString("date: " + t.Format(longform) + "\n")
		f.WriteString("tags: Blog\n====\n\n")
		fmt.Println("Your new post is ready at " + name)

	}

}

func new_page() {
	const longform = "2006/01/02 15:04:05"
	var title string
	fmt.Print("Enter the title of the page: ")
	in := bufio.NewReader(os.Stdin)
	title, _ = in.ReadString('\n')
	title = strings.TrimSpace(title)
	slug := get_slug(title)
	name := "./pages/" + slug + ".md"
	f, err := os.Create(name)
	defer f.Close()
	if err == nil {
		f.WriteString("title: " + title + "\n")
		t := time.Now()
		f.WriteString("date: " + t.Format(longform) + "\n")
		f.WriteString("tags: Blog\n====\n\n")
		fmt.Println("Your new page is ready at " + name)

	}

}
