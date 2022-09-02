package main

import (
	//"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func db_exec(sqlStatement string) {
	db, _ := sql.Open("sqlite3", AppConfig.DbPath)
	defer db.Close()
  
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal("ERROR: db_exec: ", err)
	}
}

func db_select() (links []Link) {
	db, _ := sql.Open("sqlite3", AppConfig.DbPath)
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

    	links = append(links, oneLink)
  	}

	//fmt.Println("Select all:", dbHosts)
	return links
}