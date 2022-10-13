package main

import (
	"html"
	"html/template"
	"net/http"
	"sort"
	"strings"
)

func tag(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		urlString := html.EscapeString(r.URL.Path)
		tags := strings.Split(urlString, "/")
		oneTag := tags[2]

		foundLinks := []Link{}

		AllLinks = db_select()
		split_tags()

		for _, oneLink := range AllLinks {
			if in_array(oneLink.Tags, oneTag) {
				foundLinks = append(foundLinks, oneLink)
			}
		}

		AllLinks = foundLinks
	}

	http.Redirect(w, r, "/", 302)
}

func all_tags(w http.ResponseWriter, r *http.Request) {
	type allData struct {
		Config Conf
		Tags   []string
	}
	split_tags()
	var guiData allData
	guiData.Config = AppConfig
	guiData.Tags = []string{}

	for _, oneLink := range AllLinks {
		for _, oneTag := range oneLink.Tags {
			if !in_array(guiData.Tags, oneTag) {
				guiData.Tags = append(guiData.Tags, oneTag)
			}
		}
	}

	sort.Strings(guiData.Tags)

	tmpl, _ := template.ParseFS(TemplHTML, "templates/tags.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "tags", guiData)
}

func in_array(tags []string, oneTag string) bool {
	for _, tag := range tags {
		if tag == oneTag {
			return true
		}
	}
	return false
}

func split_tags() {
	for i, oneLink := range AllLinks {
		AllLinks[i].Tags = strings.Split(oneLink.Tag, " ")
	}
}
