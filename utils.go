package main

import (
	"code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"fmt"
	"os"
	"io/ioutil"
	"crypto/sha256"
	"encoding/base64"
)

type Article struct {
	Id int
	Path string
	Hash string

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
            names = append(names, "./posts/" + f.Name())
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
				fmt.Println(stmt)
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

func main() {
	names := findfiles()
	for i := range(names) {
		hash := create_hash(names[i])
		if changed_ornot(names[i], hash) {
			fmt.Println(names[i])
		}
	}
}
