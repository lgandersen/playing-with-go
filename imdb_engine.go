package main

import (
	"fmt"
	"strconv"
	"bytes"
	"golang.org/x/net/html"
)

type MovieInfo struct {
	title string
	rating float64
	year int
	genres [3]string
}


func fast_forward(page* html.Tokenizer, n int) {
	for i := 0; i < n; i++ {
		page.Next()
	}
}


func get_rating(page* html.Tokenizer) float64 {
	fast_forward(page, 4)
	fmt.Println(string(page.Text()))
	flt, _ := strconv.ParseFloat(string(page.Text()), 64)
	return flt
}


func get_title(page* html.Tokenizer) string {
	fast_forward(page, 3)
	fmt.Println(string(page.Text()))
	return string(page.Text())
}


func get_year(page* html.Tokenizer) int {
	fast_forward(page, 4)
	fmt.Println(string(page.Text()))
	i, _ := strconv.Atoi(string(page.Text()))
	return i
}

func get_genres(page* html.Tokenizer) [3]string {
	stop := false
	idx := 0
	var genres [3]string

	fast_forward(page, 18)
	for ! stop {
		TokenType := page.Next()
		switch TokenType {
		case html.StartTagToken:
			tag_name, _ := page.TagName()
			if string(tag_name) == "a" {
				_, _, has_more_attrs := page.TagAttr()
				if has_more_attrs {
					// We have reached the date link and is thus done
					stop = true
				} else {
					page.Next()
					genres[idx] = string(page.Raw())
					idx++
					fmt.Println(string(page.Raw()))
				}
			}
		}
	}
	return genres
}


func parse_imdb_page(imdb_url string, page_raw []byte) {
	r := bytes.NewBuffer(page_raw)
	page := html.NewTokenizer(r)
	var rating float64
	var title string
	var year int
	var genres [3]string
	for {
		TokenType := page.Next()
		switch TokenType {
		case html.StartTagToken:
			_, has_attrs := page.TagName()
			if has_attrs {
				attr, value, _ := page.TagAttr()
				if string(attr) == "class" && string(value) == "ratingValue" {
					rating = get_rating(page)
				}
				if string(attr) == "class" && string(value) == "title_wrapper" {
					title = get_title(page)
					year = get_year(page)
					genres = get_genres(page)
				}
			}
		}
	}
	page_info := MovieInfo{title, rating, year, genres}
	fmt.Println(page_info)
}
