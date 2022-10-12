package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func edit_link(w http.ResponseWriter, r *http.Request) {
	type allData struct {
		Config Conf
		Link   Link
	}
	var guiData allData
	var oneLink Link

	idStr := r.FormValue("id")
	id, _ := strconv.Atoi(idStr)
	oneLink.Id = uint16(id)
	oneLink.Name = r.FormValue("name")
	oneLink.Link = r.FormValue("link")
	oneLink.Date = r.FormValue("date")
	oneLink.Tag = r.FormValue("tag")

	//fmt.Println("Edit link:", oneLink)

	guiData.Config = AppConfig
	guiData.Link = oneLink

	tmpl, _ := template.ParseFiles("templates/edit.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "edit", guiData)
}

func save_link(w http.ResponseWriter, r *http.Request) {
	var oneLink Link

	idStr := r.FormValue("id")
	id, _ := strconv.Atoi(idStr)
	oneLink.Id = uint16(id)
	oneLink.Name = r.FormValue("name")
	oneLink.Link = r.FormValue("link")
	oneLink.Date = r.FormValue("date")
	oneLink.Tag = r.FormValue("tag")

	log.Println("INFO: Edited", oneLink)

	if oneLink.Link == "" {
		fmt.Fprintf(w, "No data!")
	} else {
		db_delete(oneLink.Id)
		db_insert(oneLink)
		AllLinks = db_select()
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func del_link(w http.ResponseWriter, r *http.Request) {

	idStr := r.FormValue("id")
	id, _ := strconv.Atoi(idStr)
	dbID := uint16(id)

	log.Println("INFO: Deleted ID", dbID)

	db_delete(dbID)
	AllLinks = db_select()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
