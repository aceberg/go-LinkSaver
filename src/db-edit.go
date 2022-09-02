package main

import (
	"os"
	"log"
	"fmt"
)

func db_create() {
	if _, err := os.Stat(AppConfig.DbPath); err == nil {
        log.Println("INFO: DB exists")
    } else {
		sqlStatement := `CREATE TABLE "links" (
			"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
			"NAME"	TEXT NOT NULL,
			"LINK"	TEXT NOT NULL,
			"DATE"	TEXT NOT NULL,
			"TAG"	TEXT NOT NULL
		);`
    	db_exec(sqlStatement)
		log.Println("INFO: DB created!")
    }
}

func db_insert(oneLink Link) {
	sqlStatement := `INSERT INTO "links" (NAME, LINK, DATE, TAG)
		VALUES ('%s','%s','%s','%s');`
  	sqlStatement = fmt.Sprintf(sqlStatement, quote_str(oneLink.Name), quote_str(oneLink.Link), oneLink.Date, quote_str(oneLink.Tag))
    db_exec(sqlStatement)
}

func db_delete(id uint16) {
	sqlStatement := `DELETE FROM "links" WHERE ID='%d';`
  	sqlStatement = fmt.Sprintf(sqlStatement, id)
    db_exec(sqlStatement)
}