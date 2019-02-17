package main

import (
	"strings"
	_ "github.com/mattn/go-sqlite3"
	"regexp"
	"io/ioutil"
	"net/http"
	"time"
	"log"
)


func estimate_imdb_entry(path string) string {
	files, err := ioutil.ReadDir(path)
	checkError(err)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".nfo") {
			nfo_file, err := ioutil.ReadFile(path + "/" + file.Name())
			checkError(err)
			re := regexp.MustCompile("imdb.com/title/tt(?P<movie_id>\\d+)")
			imdb_url := re.FindString(string(nfo_file))
			return imdb_url
		}
	}
	return ""
}

type File struct {
	rowid int
	path string
	name string
}

func find_imdb_links(c Config) {
	db := open_database(c.DatabaseLocation)
	defer db.Close()
	rows, err := db.Query("SELECT rowid, path, name FROM files LIMIT 3;") // REMEMBER LIMIT
	checkError(err)

	limit := 3
	files := make([]File, 0, 1000)
	var file File
	for i := 0; i < limit; i++ {
 		rows.Next()
		err = rows.Scan(&file.rowid, &file.path, &file.name)
		files = append(files, file)
		checkError(err)
	}
	rows.Close()

	insert_stmt, err := db.Prepare("INSERT INTO imdb_pages(file_id, body) VALUES(?, ?);")
	checkError(err)
	for _, file := range files {
		imdb_url := estimate_imdb_entry(file.path + "/" + file.name)
		if imdb_url != "" {
			log.Println("Fetching url ", imdb_url)
			response, err := http.Get("https://" + imdb_url)
			checkError(err)
			body, err := ioutil.ReadAll(response.Body)
			checkError(err)
			response.Body.Close()
			log.Println("Saving contents of ", imdb_url)
			_, err = insert_stmt.Exec(file.rowid, body)
			checkError(err)
			log.Println("Inserted page from ", imdb_url)
			time.Sleep(2)
		} else {
			continue
		}
	}
	log.Println("Number of files with identified imdb-url: " + string(len(files)))
	log.Println("Identified root directory: " + c.RootDir)
}
