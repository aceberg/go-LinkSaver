package main

import (
  "net/http"
  "strings"
)

func search_links(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")

	foundLinks := []Link{}
	for _, oneLink := range AllLinks {
		if in_string(oneLink.Name, search) {
			foundLinks = append(foundLinks, oneLink)
		} else if in_string(oneLink.Link, search) {
			foundLinks = append(foundLinks, oneLink)
		} else if in_string(oneLink.Date, search) {
			foundLinks = append(foundLinks, oneLink)
		} else if in_string(oneLink.Tag, search) {
			foundLinks = append(foundLinks, oneLink)
		}
	}
	AllLinks = foundLinks

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func in_string (str1 string, str2 string) (bool) {
	return strings.Contains(
		strings.ToLower(str1),
        strings.ToLower(str2),
	)
}