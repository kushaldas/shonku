package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func new_post() {
	const longform = "2006-01-02 15:04:05.999999999 -0700 MST"
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
	const longform = "2006-01-02 15:04:05.999999999 -0700 MST"
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
		f.WriteString("====\n\n")
		fmt.Println("Your new page is ready at " + name)

	}

}

// Copies file source to destination dest.
func CopyFile(source string, dest string) (err error) {
	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, si.Mode())
		}

	}

	return
}

// Recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	fi, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return &CustomError{"Source is not a directory"}
	}

	// ensure dest dir does not already exist

	_, err = os.Open(dest)
	if !os.IsNotExist(err) {
		return &CustomError{"Destination already exists"}
	}

	// create dest dir

	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(source)

	for _, entry := range entries {

		sfp := source + "/" + entry.Name()
		dfp := dest + "/" + entry.Name()
		if entry.IsDir() {
			err = CopyDir(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		}

	}
	return
}

// A struct for returning custom error messages
type CustomError struct {
	What string
}

// Returns the error message defined in What as a string
func (e *CustomError) Error() string {
	return e.What
}
