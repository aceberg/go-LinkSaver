package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	type allData struct {
		Config Conf
		Links  []Link
	}
	split_tags()
	var guiData allData
	guiData.Config = AppConfig
	guiData.Links = AllLinks

	tmpl, _ := template.ParseFS(TemplHTML, "templates/index.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "index", guiData)
}

func home(w http.ResponseWriter, r *http.Request) {
	AllLinks = db_select()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func add_link(w http.ResponseWriter, r *http.Request) {
	var oneLink Link
	oneLink.Name = r.FormValue("name")
	oneLink.Link = r.FormValue("link")
	oneLink.Tag = r.FormValue("tag")

	if oneLink.Link == "" {
		fmt.Fprintf(w, "No data!")
	} else {
		currentTime := time.Now()
		oneLink.Date = currentTime.Format("2006-01-02")
		log.Println("INFO: Added", oneLink)
		db_insert(oneLink)
		AllLinks = db_select()
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func webgui() {
	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", index)
	http.HandleFunc("/add_link/", add_link)
	http.HandleFunc("/all_tags/", all_tags)
	http.HandleFunc("/del_link/", del_link)
	http.HandleFunc("/edit_link/", edit_link)
	http.HandleFunc("/home/", home)
	http.HandleFunc("/save_link/", save_link)
	http.HandleFunc("/search_links/", search_links)
	http.HandleFunc("/sort_links/", sort_links)
	http.HandleFunc("/tag/", tag)
	http.HandleFunc("/theme/", theme)
	http.ListenAndServe(address, nil)
}
