package main

import (
  "net/http"
  "strings"
)

func search_links(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")

	foundLinks := []Link{}
	for _, oneLink := range AllLinks {
		if strings.Contains(oneLink.Name, search) {
			foundLinks = append(foundLinks, oneLink)
		} else if strings.Contains(oneLink.Link, search) {
			foundLinks = append(foundLinks, oneLink)
		} else if strings.Contains(oneLink.Date, search) {
			foundLinks = append(foundLinks, oneLink)
		} else if strings.Contains(oneLink.Tag, search) {
			foundLinks = append(foundLinks, oneLink)
		}
	}
	AllLinks = foundLinks

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

