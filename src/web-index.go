package main

import (
  "fmt"
  "net/http"
  "html/template"
  "time"
)

func index(w http.ResponseWriter, r *http.Request) {
	type allData struct {
		Config Conf
		Links []Link
	}
	var guiData allData
	guiData.Config = AppConfig
	guiData.Links = AllLinks

	tmpl, _ := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "index", guiData)
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
		oneLink.Date = currentTime.Format("2006-01-02 15:04:05")
		fmt.Println("Link to ADD:", oneLink)
		db_insert(oneLink)
		AllLinks = db_select()
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func webgui() {
	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	fmt.Println("\n=================================== ")
	fmt.Println(fmt.Sprintf("Web GUI at http://%s", address))
	fmt.Println("=================================== \n")

	http.HandleFunc("/", index)
	http.HandleFunc("/add_link/", add_link)
	http.HandleFunc("/edit_link/", edit_link)
	http.HandleFunc("/save_link/", save_link)
	http.HandleFunc("/del_link/", del_link)
	http.HandleFunc("/sort_links/", sort_links)
	http.ListenAndServe(address, nil)
}