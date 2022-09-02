package main

import (
	"fmt"
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

func db_select(params ...string) (links []Link) {
	db, _ := sql.Open("sqlite3", AppConfig.DbPath)
	defer db.Close()

	var field, way string

	if len(params) != 2 {
		field = "DATE"
		way = "DESC"
	} else {
		field = params[0]
		way = params[1]
	}

	sqlStatement := `SELECT * FROM "links" ORDER BY %s %s`
	sqlStatement = fmt.Sprintf(sqlStatement, field, way)

	// fmt.Println("SELECT QUERY:", sqlStatement)

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