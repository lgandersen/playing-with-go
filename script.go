package main

import (
	"os"
	"fmt"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong usage. Syntax is ./script <command>")
		os.Exit(-1)
	}
	config := OpenConfiguration("./config.yaml")
	switch os.Args[1] {
	case "build_filelist_cache":
		build_filelist_cache(config)
	case "find_imdb_links":
		find_imdb_links(config)
	}
	os.Exit(0)
}
