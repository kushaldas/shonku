package main

import (
	"github.com/russross/blackfriday"
	//"fmt"
	"os"
	"io/ioutil"
)


func main() {
	md, err := ioutil.ReadFile("posts/first.md")
	if err == nil {

		output := blackfriday.MarkdownCommon(md)
		os.Stdout.Write(output)		
		//fmt.Println(s)
	}
}