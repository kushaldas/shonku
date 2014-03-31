package main

import (
	"bufio"
	"code.google.com/p/go-sqlite/go1/sqlite3"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
	"bytes"
)

type Article struct {
	Id   int
	Path string
	Hash string
}

type Post struct {
	Title 	string
	Slug 	string
	Body 	string
	Date 	string
	Tags 	[]string
}

/*
Creates the empty build database.
*/
func createdb() {
	filename := ".scrdkd.db"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// Our db is missing, time to recreate it.
		conn, err := sqlite3.Open(filename)
		if err != nil {
			fmt.Println("Unable to open the database: %s", err)
			os.Exit(1)
		}
		defer conn.Close()
		conn.Exec("CREATE TABLE builds(id INTEGER PRIMARY KEY AUTOINCREMENT, path TEXT, hash TEXT);")
	}
}

/*
Finds all the files from our posts directory.
*/
func findfiles() []string {
	files, _ := ioutil.ReadDir("./posts/")
	names := make([]string, 0)
	for _, f := range files {
		names = append(names, "./posts/"+f.Name())
	}
	return names
}

/*
Creates the hash for each files.
*/
func create_hash(filename string) string {
	md, err := ioutil.ReadFile(filename)
	if err == nil {
		data := sha256.Sum256(md)
		data2 := data[:]
		s := base64.URLEncoding.EncodeToString(data2)
		return s
		//fmt.Println(s)
	}
	return ""
}

/*
Finds out if the file content changed from the last build.
*/
func changed_ornot(filename, hash string) bool {
	db := ".scrdkd.db"
	if _, err := os.Stat(db); err == nil {
		// Our db is missing, time to recreate it.
		conn, err := sql.Open("sqlite3", db)
		if err != nil {
			fmt.Println("Unable to open the database: %s", err)
			os.Exit(1)
		}
		defer conn.Close()
		stmt := fmt.Sprintf("SELECT id, path, hash FROM builds where path='%s';", filename)
		rows, err := conn.Query(stmt)
		defer rows.Close()
		if rows.Next() {
			var article Article
			err = rows.Scan(&article.Id, &article.Path, &article.Hash)
			if article.Hash == hash {
				return false
			} else { // File hash has changed, we need to update the db
				stmt = fmt.Sprintf("UPDATE builds SET hash='%s' where id=%d;", hash, article.Id)
				rows.Close()
				_, err = conn.Exec(stmt)
				if err != nil {
					fmt.Println(err)
				}
				return true
			}
		} else { // Should be insert into DB
			conn.Exec("INSERT INTO builds(path, hash) VALUES (?, ?)", filename, hash)
			return true
		}

	}
	return true
}

/*
Returns the slug of the article
TODO:
We need to improve this function much more.
*/
func get_slug(s string) string {
	s = strings.Replace(s, "(", "-", -1)
	s = strings.Replace(s, ")", "-", -1)
	s = strings.Replace(s, " ", "-", -1)
	s = strings.Replace(s, "(", "-", -1)
	s = url.QueryEscape(s)
	return s
}

/*
Reads a post and gets all details from it.
*/
func read_post(filename string) Post {
	var buffer bytes.Buffer
	var p Post
	flag := false
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return p
	}
	defer f.Close()
	r := bufio.NewReader(f)
	titleline, err := r.ReadString('\n')
	dateline, err := r.ReadString('\n')
	line, err := r.ReadString('\n')
	tagline := line
	
	for err == nil {
		if line == "====\n" {
			flag = true
			line, err = r.ReadString('\n')
			continue
		}
		if flag {
			buffer.WriteString(line)
		}
		line, err = r.ReadString('\n')
	}


	if err == io.EOF {
		title := titleline[6:]
		title = strings.TrimSpace(title)
		date := dateline[5:]
		date = strings.TrimSpace(date)
		tagsnonstripped := strings.Split(tagline[5:], ",")
		tags := make([]string, 0)
		for i := range(tagsnonstripped) {
			word := strings.TrimSpace(tagsnonstripped[i])
			tags = append(tags, word)
		}
		
		p.Title = title
		p.Slug = get_slug(title)
		p.Body = buffer.String()
		p.Date = date
		p.Tags = tags
		
	}
	return p
}

func main() {
	names := findfiles()
	for i := range names {
		hash := create_hash(names[i])
		if changed_ornot(names[i], hash) {
			fmt.Println(names[i])
		}
	}

	s := "and but (hekko38) the9"
	fmt.Println(get_slug(s))

	read_post("./posts/first.md")
}
