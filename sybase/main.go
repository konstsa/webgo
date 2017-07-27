package main

import (
	_ "github.com/a-palchikov/sqlago"
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("sqlany", "UID=test;PWD=123456;DatabaseName=P48TestAB;Host=10.1.108.203;Port=5000")
	if err != nil {
		log.Fatalf("Unable to connect to db: %s", err)
	}

	// Run query with multiple return rows
	rows, err := db.Query("select top 20 Id,Remark from tCounter")
	if err != nil {
		log.Fatalf("Select failed: %s", err)
	}
	for rows.Next() {
		var Id int
		var Remark string
		err = rows.Scan(&Id, &Remark)
		if err != nil {
			log.Fatalf("Select failed: %s", err)
		}
		log.Printf("Id: %0.2f, Remark: %s", Id, Remark)
	}
//	defer rows.Close()
}