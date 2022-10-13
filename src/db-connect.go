package main

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

func db_exec(sqlStatement string) {
	db, _ := sql.Open("sqlite", AppConfig.DbPath)
	defer db.Close()

	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal("ERROR: db_exec: ", err)
	}
}

func db_select() (links []Link) {
	db, _ := sql.Open("sqlite", AppConfig.DbPath)
	defer db.Close()

	sqlStatement := `SELECT * FROM "links" ORDER BY DATE DESC`

	res, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal("ERROR: db_select: ", err)
	}

	links = []Link{}
	for res.Next() {
		var oneLink Link
		err = res.Scan(&oneLink.Id, &oneLink.Name, &oneLink.Link, &oneLink.Date, &oneLink.Tag)
		if err != nil {
			log.Fatal(err)
		}

		oneLink.Name = unquote_str(oneLink.Name)
		oneLink.Link = unquote_str(oneLink.Link)
		oneLink.Tag = unquote_str(oneLink.Tag)

		links = append(links, oneLink)
	}

	//fmt.Println("Select all:", dbHosts)
	return links
}
