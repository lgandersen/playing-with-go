package main


import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
)

func ignored_folders(c Config) map[string]bool {
	folders2ignore := make(map[string]bool)
	for _, folder_name := range c.FoldersToIgnore {
		folders2ignore[folder_name] = true
	}
	return folders2ignore
}

func build_filelist_cache(c Config) {
	folders2ignore := ignored_folders(c)
	db := create_database(c.DatabaseLocation)
	defer db.Close()
	insert_stmt, err := db.Prepare("INSERT INTO files(path, name) VALUES(?, ?);")
	checkError(err)
	files, err := ioutil.ReadDir(c.RootDir)
	checkError(err)
	for _, file := range files {
		if file.IsDir() == true {
			_, ignore_file := folders2ignore[file.Name()]
			if ignore_file {
				fmt.Println("Ignoring: ", file.Name())
				continue				
			}
			_, err = insert_stmt.Exec(c.RootDir, file.Name())
			checkError(err)
		}
		fmt.Printf("Path: %s, Filename: %s\n", c.RootDir, file.Name())
	}
}
