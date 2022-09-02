package main

import (
  "net/http"
  "html/template"
  "html"
  "strings"
  "sort"
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

func all_tags(w http.ResponseWriter, r *http.Request) {
	type allData struct {
		Config Conf
		Tags []string
	}
	var guiData allData

	guiData.Config = AppConfig
	guiData.Tags = []string{}

	for _, link := range AllLinks {
		if !in_array(guiData.Tags, link.Tag) {
			guiData.Tags = append(guiData.Tags, link.Tag)
		}
	}

	sort.Strings(guiData.Tags)

	tmpl, _ := template.ParseFiles("templates/tags.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "tags", guiData)
}

func in_array(tags []string, oneTag string) (bool) {
	for _, tag := range tags {
		if tag == oneTag {
			return true
		}
	}
	return false
}