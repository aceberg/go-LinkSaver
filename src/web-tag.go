package main

import (
  "net/http"
  "html"
  "strings"
)

func tag(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		urlString := html.EscapeString(r.URL.Path)
		tags := strings.Split(urlString, "/")
		oneTag := tags[2]

		foundLinks := []Link{}

		for _, oneLink := range AllLinks {
			if oneLink.Tag == oneTag {
				foundLinks = append(foundLinks, oneLink)
			} 
		}

		AllLinks = foundLinks
	} 

	http.Redirect(w, r, "/", http.StatusSeeOther)
}