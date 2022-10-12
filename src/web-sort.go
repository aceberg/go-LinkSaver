package main

import (
	"net/http"
	"sort"
)

func sort_links(w http.ResponseWriter, r *http.Request) {

	sort_method := r.FormValue("sort_method")

	switch sort_method {
	case "name-up":
		sort.SliceStable(AllLinks, func(i, j int) bool {
			return AllLinks[i].Name < AllLinks[j].Name
		})
	case "name-down":
		sort.SliceStable(AllLinks, func(i, j int) bool {
			return AllLinks[i].Name > AllLinks[j].Name
		})
	case "link-up":
		sort.SliceStable(AllLinks, func(i, j int) bool {
			return AllLinks[i].Link < AllLinks[j].Link
		})
	case "link-down":
		sort.SliceStable(AllLinks, func(i, j int) bool {
			return AllLinks[i].Link > AllLinks[j].Link
		})
	case "date-up":
		sort.SliceStable(AllLinks, func(i, j int) bool {
			return AllLinks[i].Date < AllLinks[j].Date
		})
	case "date-down":
		sort.SliceStable(AllLinks, func(i, j int) bool {
			return AllLinks[i].Date > AllLinks[j].Date
		})
	case "tag-up":
		sort.SliceStable(AllLinks, func(i, j int) bool {
			return AllLinks[i].Tag < AllLinks[j].Tag
		})
	case "tag-down":
		sort.SliceStable(AllLinks, func(i, j int) bool {
			return AllLinks[i].Tag > AllLinks[j].Tag
		})
	default:
		AllLinks = db_select()
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
