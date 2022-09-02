package main

import (
  "net/http"
)

func sort_links(w http.ResponseWriter, r *http.Request) {

	sort_method := r.FormValue("sort_method")

	switch sort_method {
	case "name-up":
		AllLinks = db_select("NAME","ASC")
	case "name-down":
		AllLinks = db_select("NAME","DESC")
	case "link-up":
		AllLinks = db_select("LINK","ASC")
	case "link-down":
		AllLinks = db_select("LINK","DESC")
	case "date-up":
		AllLinks = db_select("DATE","ASC")
	case "date-down":
		AllLinks = db_select("DATE","DESC")
	case "tag-up":
		AllLinks = db_select("TAG","ASC")
	case "tag-down":
		AllLinks = db_select("TAG","DESC")
	default:
		AllLinks = db_select()
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}